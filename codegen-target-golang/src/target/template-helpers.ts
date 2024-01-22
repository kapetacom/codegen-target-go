/**
 * Copyright 2023 Kapeta Inc.
 * SPDX-License-Identifier: MIT
 */
import Handlebars = require('handlebars');
import {parseEntities, Template} from '@kapeta/codegen-target';
import {HelperOptions} from 'handlebars';
import {parseKapetaUri} from '@kapeta/nodejs-utils';
import {
    asComplexType,
    DataTypeReader,
    DSLData, DSLDataType, DSLDataTypeProperty,
    DSLEntity,
    DSLEntityType,
    DSLReferenceResolver,
    DSLType,
    GoWriter,
    RESTMethodParameterReader,
    RESTMethodReader,
    ucFirst,
} from '@kapeta/kaplang-core';

const DB_TYPES = ['kapeta/resource-type-mongodb', 'kapeta/resource-type-postgresql'];
export type HandleBarsType = typeof Handlebars;

const githubLocation = "github.com/kapeta/todo";
export const addTemplateHelpers = (engine: HandleBarsType, data: any, context: any): void => {
    const TypeMap: { [key: string]: string } = {
        Instance: 'InstanceValue',
        InstanceProvider: 'InstanceProviderValue',
    };

    let parsedEntities: DSLData[] | undefined = undefined;

    function getParsedEntities(): DSLData[] {
        if (!parsedEntities && context.spec?.entities?.source?.value) {
            parsedEntities = parseEntities(context.spec?.entities?.source?.value);
        }

        if (!parsedEntities) {
            return [];
        }

        return parsedEntities as DSLData[];
    }

    const resolvePath = (path: string, options: HelperOptions) => {
        let fullPath = path;
        if (options.hash.base) {
            let baseUrl: string = options.hash.base;
            while (baseUrl.endsWith('/')) {
                baseUrl = baseUrl.substring(0, baseUrl.length - 1);
            }
            if (!fullPath.startsWith('/')) {
                fullPath = '/' + fullPath;
            }

            fullPath = baseUrl + fullPath;
        }
        return fullPath;
    };

    engine.registerHelper('valueType', (value: DSLType) => {
        const type = asComplexType(value);
        if (TypeMap[type.name]) {
            type.name = TypeMap[type.name];
        }
        return Template.SafeString(GoWriter.toTypeCode(type));
    });

    engine.registerHelper('returnType', (value: DSLType) => {
        if (!value) {
            return 'void';
        }

        return Template.SafeString(GoWriter.toTypeCode(value));
    });

    engine.registerHelper('path', resolvePath);

    engine.registerHelper('consumes-databases', function (this: any, options) {
        const consumers = context.spec.consumers;

        if (!consumers || consumers.length === 0) {
            return '';
        }

        if (
            consumers.some((consumer: any) => {
                return DB_TYPES.some((dbType) => consumer.kind.includes(dbType));
            })
        ) {
            return options.fn(this);
        }

        return '';
    });

    engine.registerHelper('echoURLPath', (path, options: HelperOptions) => {
        const fullPath = resolvePath(path, options);

        return fullPath.replace(/\{([^}]+)}/g, ':$1').toLowerCase();
    });

    const $toTypeMap = (method: RESTMethodReader, transport: string) => {
        if (!method.parameters) {
            return Template.SafeString('void');
        }

        const transportArgs: RESTMethodParameterReader[] = method.parameters.filter(
            (value: RESTMethodParameterReader) =>
                value.transport && value.transport.toLowerCase() === transport.toLowerCase()
        );

        if (transportArgs.length === 0) {
            return Template.SafeString('void');
        }

        return Template.SafeString(
            '{' +
            transportArgs
                .map(
                    (value) =>
                        `'${value.name}'${value.optional ? '?' : ''}: ${GoWriter.toTypeCode(value.type)}`
                )
                .join(', ') +
            '}'
        );
    };

    engine.registerHelper('paramsMap', (method: RESTMethodReader) => {
        return $toTypeMap(method, 'path');
    });

    engine.registerHelper('queryMap', (method: RESTMethodReader) => {
        return $toTypeMap(method, 'query');
    });

    engine.registerHelper('bodyType', (method: RESTMethodReader) => {
        if (!method.parameters) {
            return Template.SafeString('void');
        }

        const bodyArgument = method.parameters.find(
            (value) => value.transport && value.transport.toLowerCase() === 'body'
        );

        if (!bodyArgument) {
            return Template.SafeString('void');
        }

        return GoWriter.toTypeCode(bodyArgument.type);
    });

    engine.registerHelper('fullName', (value: string) => {
        const uri = parseKapetaUri(value);
        return uri.fullName;
    });

    engine.registerHelper('golang-imports-dto', function (arg: DSLEntity) {
        const entities = getParsedEntities();
        const resolver = new DSLReferenceResolver();
        const referencesEntities = resolver.resolveReferencesFrom([arg], entities);

        if (referencesEntities.length === 0) {
            return '';
        }
        const result = referencesEntities
            .map((entity) => {
                const native = DataTypeReader.getNative(entity);
                if (native) {
                    return '';
                }
                return ''; //`import ${githubLocation}/entities"`;
            }).join('\n')

        if (arg.type === DSLEntityType.DATATYPE) {
            // cast to DSLEntityType.DATATYPE
            const datatype = arg as DSLDataType;
            const imports = datatype.properties?.map((property) => {
             const thisProp = property.type as DSLDataTypeProperty;
                if (thisProp.name === "date") {
                    return  Template.SafeString(`import "time"`);
                }
                return '';
            }).filter(item => item !== "");

            if (imports && imports.length > 0) {
                return Template.SafeString(imports.join('\n'));
            }
            return '';
        }
        return Template.SafeString(result);
    });

    engine.registerHelper('typescript-imports-config', function (arg: DSLEntity) {
        const entities = getParsedEntities();
        const resolver = new DSLReferenceResolver();
        const referencesEntities = resolver.resolveReferencesFrom([arg], entities);

        if (referencesEntities.length === 0) {
            return '';
        }

        return Template.SafeString(
            referencesEntities
                .map((entity) => {
                    const native = DataTypeReader.getNative(entity);
                    if (native) {
                        return `import { ${entity.name} } from "${native}";`;
                    }

                    return `import { ${ucFirst(entity.name)}Config } from './${ucFirst(entity.name)}';`;
                })
                .join('\n')
        );
    });

    engine.registerHelper('golang-dto', (entity: DSLData) => {
        const writer = new GoWriter();

        try {
            return Template.SafeString(writer.write([entity]));
        } catch (e) {
            console.warn('Failed to write entity', entity);
            throw e;
        }
    });

    engine.registerHelper('typescript-config', (entity: DSLData) => {
        const writer = new GoWriter();

        try {
            // All config entities are postfixed with Config
            const copy = {...entity, name: entity.name + 'Config'};
            return Template.SafeString(writer.write([copy]));
        } catch (e) {
            console.warn('Failed to write entity', entity);
            throw e;
        }
    });
};
