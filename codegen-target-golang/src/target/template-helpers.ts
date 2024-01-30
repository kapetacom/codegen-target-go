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

    // return type for interfaces this is either error or the entity and error
    engine.registerHelper('returnTypeInterface', (value: DSLType) => {
        const type = asComplexType(value);
        if(type.name === 'void') {
            return Template.SafeString('error');
        }
        const entities = getParsedEntities();
        for (const entity of entities) {
            if (entity.name === type.name) {
                return Template.SafeString("*entities."+GoWriter.toTypeCode(value)+", error");
            }
        }
        const returnValue = GoWriter.toTypeCode(value)
        if(returnValue === '') {
            return Template.SafeString('error');
        }
        return Template.SafeString(GoWriter.toTypeCode(value)+", error");
    
    });

    // returns the variable type for the return type including package name
    engine.registerHelper('variableType', (value: DSLType) => {
        const type = asComplexType(value);
        const entities = getParsedEntities();
        for (const entity of entities) {
            if (entity.name === type.name) {
                return Template.SafeString("*entities."+GoWriter.toTypeCode(value));
            }
        }
        return Template.SafeString(GoWriter.toTypeCode(value));
    });

    engine.registerHelper('hasReturnValue', (value: DSLType) => {
        const type = asComplexType(value);
        if(type.name === 'void') {
            return false;
        }
        return true;
    });

    engine.registerHelper('returnType', (value: DSLType) => {
        if (!value) {
            return '';
        }
        const entities = getParsedEntities();

        if(entities.length === 0) {
            return Template.SafeString(GoWriter.toTypeCode(value));
        }
        const type = asComplexType(value);
        if(type.name === 'void') {
            return Template.SafeString('');
        }

        for (const entity of entities) {
            if (entity.name === type.name) {
                return Template.SafeString("*entities."+GoWriter.toTypeCode(value));
            }
        }
        return Template.SafeString(GoWriter.toTypeCode(value));
    });
    
    engine.registerHelper('returnTypeDefaultValue', (value: DSLType) => {
        if (!value) {
            return '';
        }
        const entities = getParsedEntities();

        if(entities.length === 0) {
            return Template.SafeString(GoWriter.toTypeCode(value));
        }
        const type = asComplexType(value);
        if(type.name === 'void') {
            return Template.SafeString('nil');
        }

        for (const entity of entities) {
            if (entity.name === type.name) {
                return Template.SafeString("entities."+GoWriter.toTypeCode(value)+"{}");
            }
        }
        const gotype = GoWriter.toTypeCode(value)
        switch(gotype) {
            case "string":
                return Template.SafeString('""');
            case "int":
                return Template.SafeString('0');
            case "bool":
                return Template.SafeString('false');
            default:
                return Template.SafeString('nil');
        }
    });

    engine.registerHelper('ucFirst', (value: string) => {
        return ucFirst(value);
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

    engine.registerHelper('urlString', (path, options: HelperOptions) => {
        // take the url and replace all parameters with fmt.Sprintf("%v", param)
        const fullPath = resolvePath(path, options);
        return fullPath.replace(/\{([^}]+)}/g, '%v').toLowerCase();
    });

    const $toTypeMap = (method: RESTMethodReader, transport: string) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const transportArgs: RESTMethodParameterReader[] = method.parameters.filter(
            (value: RESTMethodParameterReader) =>
                value.transport && value.transport.toLowerCase() === transport.toLowerCase()
        );

        if (transportArgs.length === 0) {
            return Template.SafeString('');
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

    const $toTypeList = (method: RESTMethodReader, transport: string) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const transportArgs: RESTMethodParameterReader[] = method.parameters.filter(
            (value: RESTMethodParameterReader) =>
                value.transport && value.transport.toLowerCase() === transport.toLowerCase()
        );

        if (transportArgs.length === 0) {
            return Template.SafeString('');
        }

        return Template.SafeString(","+
            transportArgs
                .map(
                    (value) =>
                        `${value.name}`
                )
                .join(', ')
        );
    };

    // returns a list of the variables for the @path parameters in the order they appear in the path
    engine.registerHelper('pathParamsList', (method: RESTMethodReader) => {
        return $toTypeList(method, 'path');
    });

    // returns a list of the variables for the @body parameters in the order they appear in the path
    engine.registerHelper('bodyParamsList', (method: RESTMethodReader) => {
        return $toTypeList(method, 'body');
    });

    // returns a list of the variables for the @path parameters in the order they appear in the path
    engine.registerHelper('queryParamsList', (method: RESTMethodReader) => {
        return $toTypeList(method, 'query');
    });
    
    engine.registerHelper('paramsMap', (method: RESTMethodReader) => {
        return $toTypeMap(method, 'path');
    });

    engine.registerHelper('queryMap', (method: RESTMethodReader) => {
        return $toTypeMap(method, 'query');
    });

    engine.registerHelper('bodyType', (method: RESTMethodReader) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const bodyArgument = method.parameters.find(
            (value) => value.transport && value.transport.toLowerCase() === 'body'
        );

        if (!bodyArgument) {
            return Template.SafeString('');
        }

        return GoWriter.toTypeCode(bodyArgument.type);
    });

    engine.registerHelper('queryParametersVariables', (method: RESTMethodReader) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const queryParameters = method.parameters.filter(
            (value) => value.transport && value.transport.toLowerCase() === 'query'
        );

        if (queryParameters.length === 0) {
            return Template.SafeString('');
        }

        return Template.SafeString(
            queryParameters
                .map((value) => {
                    return `_ = ctx.Param("${value.name}")`;
                })
                .join('\n')
        );
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

    engine.registerHelper('golang-import-entities',function (value) {
        console.log("golang-import-entities...........................")
        const entities = getParsedEntities();
        if (entities.length === 0) {
            return false;
        }
        return true;
    });
    engine.registerHelper('go-imports-config', function (arg: DSLEntity) {
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

    engine.registerHelper('go-config', (entity: DSLData) => {
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
