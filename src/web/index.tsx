/**
 * Copyright 2023 Kapeta Inc.
 * SPDX-License-Identifier: MIT
 */
import React from 'react';

import { ILanguageTargetProvider } from '@kapeta/ui-web-types';
import { FormField } from '@kapeta/ui-web-components';

// @ts-ignore
import kapetaDefinition from '../../kapeta.yml';
// @ts-ignore
import packageJson from '../../package.json';

interface GoTargetConfigOptions {
    basePackage: string;
}

function validatePackageName(fieldName: string, value: string) {
    const goPackagePathRegex = /^([a-z0-9_\-.]+)(?::([0-9]+))?(\/[^?\s]*)?/;
    if (!goPackagePathRegex.test(value)) {
        throw new Error('Not a valid Go package path');
    }
}

const GoTargetConfig = () => {
    return (
        <>
            <FormField
                name={'spec.target.options.basePackage'}
                label={'Package name'}
                validation={['required', validatePackageName]}
                help={'Must be a valid Go package name. E.g. github.com/kapetacom/example'}
            />
        </>
    );
};
const targetConfig: ILanguageTargetProvider<GoTargetConfigOptions> = {
    kind: kapetaDefinition.metadata.name,
    version: packageJson.version,
    title: kapetaDefinition.metadata.title,
    blockKinds: ['kapeta/block-type-service', 'kapeta/block-type-mcp'],
    resourceKinds: [
        'kapeta/resource-type-rest-api',
        'kapeta/resource-type-rest-client',
        'kapeta/resource-type-auth-jwt-consumer',
        'kapeta/resource-type-external-services',
        'kapeta/resource-type-rabbitmq-subscriber',
        'kapeta/resource-type-rabbitmq-publisher',
        'kapeta/resource-type-mongodb',
        'kapeta/resource-type-pubsub-publisher',
        'kapeta/resource-type-pubsub-subscriber',
        'kapeta/resource-type-mcp-tools-server',
        'kapeta/resource-type-mcp-tools-client',
        //        'kapeta/resource-type-postgresql',
        //        'kapeta/resource-type-redis',
        //        'kapeta/resource-type-smtp-client',
        //        'kapeta/resource-type-auth-jwt-provider',
        //        'kapeta/resource-type-cloud-bucket',
    ],
    editorComponent: GoTargetConfig,
    definition: kapetaDefinition,
    validate: (options: any) => {
        const errors: string[] = [];

        if (!options) {
            errors.push('Missing target configuration');
        } else {
            if (!options.basePackage) {
                errors.push('Missing base package');
            } else {
                try {
                    validatePackageName('basePackage', options.basePackage);
                } catch (e) {
                    errors.push('Base package is invalid');
                }
            }
        }
        return errors;
    },
};

export default targetConfig;
