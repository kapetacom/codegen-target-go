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

    function fixType(type: DSLType): string {
        const localType = asComplexType(type);
        const entities = getParsedEntities();
        for (const entity of entities) {
            if (entity.name === localType.name) {
                return `*entities.${GoWriter.toTypeCode(type)}`;
            }
        }
        if (localType.list) {
            return `*${GoWriter.toTypeCode(type)}`;
        }
        if (localType.name === "Set") {
            return `*${GoWriter.toTypeCode(type)}`;
        }
        const res = GoWriter.toTypeCode(type)
        if (res.startsWith("map")) {
            const type = res.split("]")[1]
            if (!isBuiltInType(type)) {
                return res.split("]")[0] + "]*entities." + type
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
        const entities = getParsedEntities();
        for (const entity of entities) {
            if (entity.name === type.name) {
                return Template.SafeString("*entities." + GoWriter.toTypeCode(value) + ", error");
            }
        }
        const returnValue = GoWriter.toTypeCode(value)
        if (returnValue === '') {
            return Template.SafeString('error');
        }
        return Template.SafeString(GoWriter.toTypeCode(value) + ", error");

    });

    // returns the variable type for the return type including package name
    engine.registerHelper('variableType', (value: DSLType) => {
        return Template.SafeString(fixType(value));
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
        const entities = getParsedEntities();

        if (entities.length === 0) {
            return Template.SafeString(GoWriter.toTypeCode(value));
        }
        const type = asComplexType(value);
        if (type.name === 'void') {
            return Template.SafeString('');
        }
        return Template.SafeString(fixType(value));
    });

    engine.registerHelper('returnTypeDefaultValue', (value: DSLType) => {
        if (!value) {
            return '';
        }
        const entities = getParsedEntities();

        if (entities.length === 0) {
            if (GoWriter.toTypeCode(value) === 'void') {
                return Template.SafeString('nil');
            }
            return Template.SafeString(GoWriter.toTypeCode(value));
        }
        const type = asComplexType(value);
        if (type.name === 'void') {
            return Template.SafeString('nil');
        }

        for (const entity of entities) {
            if (entity.name === type.name) {
                return Template.SafeString("entities." + GoWriter.toTypeCode(value) + "{}");
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


    engine.registerHelper('bodyParametersVariables', (method: RESTMethodReader) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const queryParameters = method.parameters.filter(
            (value) => value.transport && value.transport.toLowerCase() === 'body'
        );

        if (queryParameters.length === 0) {
            return Template.SafeString('');
        }
        const entities = getParsedEntities();
        return Template.SafeString(
            queryParameters
                .map((value) => {
                    let typename = GoWriter.toTypeCode(value.type);
                    if (!isBuiltInType(value.type)) {
                        typename = `&entities.${GoWriter.toTypeCode(value.type)}`;
                    }
                    let valueName = value.name
                    if (valueName === "type") {
                        valueName = "_type"
                    }
                    let out = `${valueName} := ${typename}{}\n`
                    out += `if err = request.GetBody(ctx, &${valueName}); err != nil {\n`
                    out += `return ctx.String(400, fmt.Sprintf("bad request, unable to unmarshal ${value.name} %v", err))\n}`;
                    return out;
                })
                .join('\n')
        );
    });


    engine.registerHelper('pathParametersVariables', (method: RESTMethodReader) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const queryParameters = method.parameters.filter(
            (value) => value.transport && value.transport.toLowerCase() === 'path'
        );

        if (queryParameters.length === 0) {
            return Template.SafeString('');
        }
        const entities = getParsedEntities();
        return Template.SafeString(
            queryParameters
                .map((value) => {
                    let typename = `${GoWriter.toTypeCode(value.type)}`;
                    for (const entity of entities) {
                        if (entity.name === value.type.name) {
                            typename = `entities.${GoWriter.toTypeCode(value.type)}`;
                        }
                    }
                    let valueName = value.name
                    if (valueName === "type") {
                        valueName = "_type"
                    }
                    let out = `var ${valueName} ${typename}\n`
                    out += `if err = request.GetPathParams(ctx, "${value.name}", &${valueName}); err != nil {\n`
                    out += `return ctx.String(400, fmt.Sprintf("bad request, unable to get path param ${value.name} %v", err))\n}`;
                    return out
                })
                .join('\n')
        );
    });

    engine.registerHelper('headerParametersVariables', (method: RESTMethodReader) => {
        if (!method.parameters) {
            return Template.SafeString('');
        }

        const queryParameters = method.parameters.filter(
            (value) => value.transport && value.transport.toLowerCase() === 'header'
        );

        if (queryParameters.length === 0) {
            return Template.SafeString('');
        }
        const entities = getParsedEntities();
        return Template.SafeString(
            queryParameters
                .map((value) => {
                    let typename = fixType(value.type)
                    if(!typename.startsWith("*")){
                        typename = "*" + typename
                    }
                    let valueName = value.name
                    if (valueName === "type") {
                        valueName = "_type"
                    }
                    let out = `var ${valueName} ${typename}\n`
                    out += `if err = request.GetHeaderParams(ctx, "${value.name}", ${valueName}); err != nil {\n`
                    out += `return ctx.String(400, fmt.Sprintf("bad request, unable to get path param ${value.name} %v", err))\n}`;
                    return out
                })
                .join('\n')
        );
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
        const entities = getParsedEntities();
        return Template.SafeString(
            queryParameters
                .map((value) => {
                    const typename = fixType(value.type);

                    let valueName = value.name
                    if (valueName === "type") {
                        valueName = "_type"
                    }

                    let out = `var ${valueName} ${typename}\n`
                    out += `if err = request.GetQueryParam(ctx, "${value.name}", ${valueName}); err != nil {\n`
                    out += `return ctx.String(400, fmt.Sprintf("bad request, unable to get query param ${value.name} %v", err))\n}`;
                    return out
                })
                .join('\n')
        );
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
                let valueName = value.name
                if (valueName === "type") {
                    valueName = "_type"
                }
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

                let valueName = value.name
                if (valueName === "type") {
                    valueName = "_type"
                }
                if (!isBuiltInType(value.type)) {
                    return Template.SafeString(`${valueName}`)
                }
                return valueName;

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
                if (includeTypes) {
                    let valueName = value.name
                    if (valueName === "type") {
                        valueName = "_type"
                    }

                    return Template.SafeString(`${valueName} ${fixType(value.type)}`);
                }
                let valueName = value.name
                if (valueName === "type") {
                    valueName = "_type"
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

        if (referencesEntities.length === 0) {
            return '';
        }
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
            }).filter(item => item !== "");

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
