/**
 * Copyright 2023 Kapeta Inc.
 * SPDX-License-Identifier: MIT
 */

import {format, Target, SourceFile, GeneratedFile} from '@kapeta/codegen-target';
import type {GeneratedAsset} from '@kapeta/codegen';
import Path from 'path';
import {execSync} from 'child_process';
import {mergeDevcontainers} from './target/merge-devcontainers';
import {addTemplateHelpers} from './target/template-helpers';

export default class GoTarget extends Target {
    constructor(options: any) {
        super(options, Path.resolve(__dirname, '../'));
    }

    mergeFile(sourceFile: SourceFile, targetFile: GeneratedFile, lastFile: GeneratedFile | null) {
        if (sourceFile.filename === '.devcontainer/devcontainer.json') {
            return mergeDevcontainers(sourceFile, targetFile, lastFile);
        }

        return super.mergeFile(sourceFile, targetFile, lastFile);
    }

    protected _createTemplateEngine(data: any, context: any) {
        const engine = super._createTemplateEngine(data, context);

        addTemplateHelpers(engine, data, context);

        return engine;
    }

    protected _postProcessCode(filename: string, code: string) {
        if (filename.endsWith('.go')) {
            try {
                code = execSync(`gofmt `, {input: code, encoding: 'utf-8'});
            } catch (error: any) {
                // Handle errors
                console.log(`There were an error in the following code ${code} and the error is ${error}`);
                throw new Error(`Error executing gofmt/goimports: ${error.stderr || error.message}`);
            }
        }
        return format(filename, code);
    }

    generate(data: any, context: any): GeneratedFile[] {
        return super.generate(data, context);
    }

    //This is only execute via the kap cli
    async postprocess(targetDir: string, files: GeneratedAsset[]): Promise<void> {
        for (const file of files) {
            if (file.filename.endsWith('.go')) {
                console.log('Running gofmt on %s', file.filename);
                execSync(`gofmt -w ${file.filename}`);
            }
        }

        try {
            console.log('Running go mod tidy in %s', targetDir);
            execSync('go mod tidy', {
                cwd: targetDir,
            });
            console.log('done tidying');
        } catch (error: any) {
            // Handle errors
            throw new Error(`Error executing go mod tidy: ${error.stderr || error.message}`);
        }

    }
}
