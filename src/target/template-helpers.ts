/**
 * Copyright 2023 Kapeta Inc.
 * SPDX-License-Identifier: MIT
 */
import Handlebars = require('handlebars');
import {parseEntities, Template} from '@kapeta/codegen-target';
import {HelperOptions} from 'handlebars';
import {
    asComplexType,
    DataTypeReader,
    DSLData, DSLDataType, DSLDataTypeProperty,
    DSLEntity,
    DSLEntityType, DSLMethod,
    DSLReferenceResolver,
    DSLType, DSLTypeHelper,
    GoWriter,
    RESTMethodParameterReader,
    RESTMethodReader,
    ucFirst,
} from '@kapeta/kaplang-core';
import {DSLController} from "@kapeta/kaplang-core/src/interfaces";
import isBuiltInType = DSLTypeHelper.isBuiltInType;
import { get } from 'lodash';


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

    function getEntityType(type: DSLType): DSLData | undefined {
        const localType = asComplexType(type);
        const entities = getParsedEntities();
        for (const entity of entities) {
            if (entity.name === localType.name) {
                return entity;
            }
        }
        return undefined;
    }

    /**
     * Wraps GoWriter.toTypeCode to handle naming the package of the types
     */
    function toGoTypeCode(type: DSLType, prefix = ''): string {
        const localType = asComplexType(type);

        const entityType = getEntityType(type);
        if (entityType) {
            if (localType.list) {
                return GoWriter.toTypeCode(type).replace(/^\[]/, `[]${prefix}entities.`);
            }
            return `${prefix}entities.${GoWriter.toTypeCode(type)}`;
        }

        if (localType.list) {
            return `${prefix}${GoWriter.toTypeCode(type)}`;
        }
        if (localType.name === "Set") {
            return `${prefix}${GoWriter.toTypeCode(type)}`;
        }
        const res = GoWriter.toTypeCode(type)
        if (res.startsWith("map[")) {
            const type = res.split("]")[1]
            if (!isBuiltInType(type)) {
                return `${res.split("]")[0]}]${prefix}entities.${type}`
            }
        }

        return res;
    }

    // return type for interfaces this is either error or the entity and error
    engine.registerHelper('returnTypeInterface', (value: DSLType) => {
        const type = asComplexType(value);
        if (type.name === 'void') {
            return Template.SafeString('error');
        }

        const returnValue = GoWriter.toTypeCode(value)
        if (returnValue === '') {
            return Template.SafeString('error');
        }
        return Template.SafeString(toGoTypeCode(value, '*') + ", error");

    });

    engine.registerHelper('isEntityType', (type: DSLType, options: HelperOptions) => {
        const entityType = getEntityType(type);
        if (entityType) {
            return options.fn(this);
        }
        return options.inverse(this);
    });

    // returns the variable type for the return type including package name
    engine.registerHelper('variableType', (value: DSLType, options: HelperOptions) => {
        return Template.SafeString(toGoTypeCode(value, options.hash['prefix'] ?? '*'));
    });

    engine.registerHelper('hasReturnValue', (value: DSLType) => {
        const type = asComplexType(value);
        if (type.name === 'void') {
            return false;
        }
        return true;
    });

    engine.registerHelper('returnType', (value: DSLType) => {
        if (!value) {
            return '';
        }

        const type = asComplexType(value);
        if (type.name === 'void') {
            return Template.SafeString('');
        }
        return Template.SafeString(toGoTypeCode(value, '*'));
    });

    engine.registerHelper('returnTypeDefaultValue', (value: DSLType) => {
        if (!value) {
            return '';
        }
        const type = asComplexType(value);
        if (type.name === 'void') {
            return Template.SafeString('nil');
        }

        const entities = getParsedEntities()
        for (const entity of entities) {
            if (entity.name === type.name) {
                return Template.SafeString(toGoTypeCode(value, '') + "{}");
            }
        }

        const gotype = GoWriter.toTypeCode(value)
        switch (gotype) {
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

        return Template.SafeString("," +
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

    engine.registerHelper('generateRequestParameters', (method: RESTMethodReader) => {
        let out = 'type RequestParameters  struct {\n'
        if (method.parameters) {
            method.parameters.forEach((value) => {
                let typename = toGoTypeCode(value.type, '*');
                const valueName = value.name

                let required = value.optional ? "" : ";required"
                if(value.transport === "BODY") {
                    // we can't not have a required body, since the definition of required in httpin is that the 'field' is there
                    // and a body on a POST is always there.
                    required = ""
                }
                const transport = value.transport === "BODY" ? "body=json" : `${value.transport.toLowerCase()}=${value.name}`
                out += `${ucFirst(valueName)} ${typename} \`in:"${transport}${required}"\`\n`
            });
        }
        out += '}'
        return Template.SafeString(out);
    });


    engine.registerHelper('queryParametersFunctions', (method: RESTMethodReader) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const queryParameters = method.parameters.filter(
            (value) => value.transport && value.transport.toLowerCase() === 'query'
        );

        if (queryParameters.length === 0) {
            return Template.SafeString('');
        }
        let result = queryParameters
            .map((value) => {
                const valueName = value.name
                let out = `client.QueryParameterRequestModifier(${valueName})`;
                return out
            })
            .join(',')
        if (result !== "") {
            return Template.SafeString("," + result);
        }
        return result;
    });

    engine.registerHelper('parametersNeedError', (method: RESTMethodReader, options: HelperOptions) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const parameterThatNeedsError = method.parameters.some(
            (value) => value.transport && ['query', 'body', 'path'].includes(value.transport.toLowerCase())
        );

        if (parameterThatNeedsError) {
            return options.fn(this);
        }
        return options.inverse(this);
    });

    engine.registerHelper('service_interface_arguments', (method: RESTMethodReader) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const parameters = method.parameters.filter(
            (value) => value.transport
        );

        if (parameters.length === 0) {
            return Template.SafeString('');
        }
        let args = parameters
            .map((value) => {
                const valueName = value.name
                return `params.${ucFirst(valueName)}`
            })
            .join(' ,')
        // prefix args with , if not there
        if (args.length > 0 && !args.startsWith(",")) {
            args = ", " + args;
        }
        return Template.SafeString(args);
    });

    function getRestParameters(method: RESTMethodReader, includeTypes: boolean) {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const parameters = method.parameters.filter(
            (value) => value.transport
        );

        if (parameters.length === 0) {
            return Template.SafeString('');
        }
        let args = parameters
            .map((value) => {
                const valueName = GoWriter.restMethodParameterReaderReservedWordMapper(value)
                if (includeTypes) {
                    return Template.SafeString(`${valueName} ${toGoTypeCode(value.type, '*')}`);
                }
                return valueName;

            })
            .join(' ,')
        // prefix args with , if not there
        if (args.length > 0 && !args.startsWith(",")) {
            args = ", " + args;
        }
        return args;
    }

    engine.registerHelper('requestparameterarguments', (method: RESTMethodReader) => {
        return getRestParameters(method, false);
    });
    engine.registerHelper('request_parameter_arguments_with_type', (method: RESTMethodReader) => {
        return getRestParameters(method, true);
    });

    engine.registerHelper('golang-imports-dto', function (arg: DSLEntity) {
        const entities = getParsedEntities();
        const resolver = new DSLReferenceResolver();
        const referencesEntities = resolver.resolveReferencesFrom([arg], entities);

        const result = referencesEntities
            .map((entity) => {
                const native = DataTypeReader.getNative(entity);
                if (native) {
                    const [packageName] = entity.name.split('.')
                    return `import ${packageName} "${native}"`;
                }
                return ''; //`import ${githubLocation}/entities"`;
            }).join('\n')

        if (arg.type === DSLEntityType.DATATYPE) {
            // cast to DSLEntityType.DATATYPE
            const datatype = arg as DSLDataType;
            const imports = datatype.properties?.map((property) => {
                const thisProp = property.type as DSLDataTypeProperty;
                if (thisProp.name === "date") {
                    return Template.SafeString(`import kapeta "github.com/kapetacom/sdk-go-config"`);
                }
                return '';
            }).filter(item => item !== "")
            .filter((value, index, self) => {
                const uniqueSet = new Set(self.map(item => item.toString()));
                return uniqueSet.has(value.toString()) && [...uniqueSet].indexOf(value.toString()) === index;
            });

            if (imports && imports.length > 0) {
                return Template.SafeString(imports.join('\n'));
            }
            return '';
        }

        return Template.SafeString(result);
    });

    // returns true if the controller has a method that returns non built in type
    // this excludes a check for the return type
    engine.registerHelper('golang-import-entities', function (arg: DSLController) {
        const entities = getParsedEntities();
        if (entities.length === 0) {
            return false;
        }
        for (const method of arg.methods) {

            for (const parameter of method.parameters ?? []) {
                const type = parameter.type;
                const is = isBuiltInType(type);
                if (!is) {
                    return true;
                }

            }
        }

        return false;
    });

    // returns true if the controller has a method that returns non built in type
    // this cheeks both the return type and the parameters of the method
    engine.registerHelper('golang-import-entities-including-returntype', function (arg: DSLController) {
        const entities = getParsedEntities();
        if (entities.length === 0) {
            return false;
        }
        for (const method of arg.methods) {
            if (method.returnType) {
                const complexType = asComplexType(method.returnType);
                for (const entity of entities) {
                    if (entity.name === complexType.name) {
                        return true;
                    }
                }
            }


            for (const parameter of method.parameters ?? []) {
                const type = parameter.type;
                const complexType = asComplexType(type);
                for (const entity of entities) {
                    if (entity.name === complexType.name) {
                        return true;
                    }
                }
            }
        }

        return false;
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
                        const [packageName] = entity.name.split('.')
                        return `import ${packageName} "${native}"`;
                    }
                })
                .join('\n')
        );
    });

    engine.registerHelper('golang-dto', (entity: DSLData) => {
        const writer = new GoWriter();

        try {
           // entity.name = getName(entity.name)
            let out = writer.write([entity]);
            if (entity.type === DSLEntityType.ENUM) {
                const name = entity.name;
                out += `\nfunc (s *${name}) ToString() (string, error) {
                    return string(*s), nil
                }
                func (s *${name}) FromString(x string) error {
                    *s = ${name}(x)
                    return nil
                }\n`
            }
            return Template.SafeString(out);
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
