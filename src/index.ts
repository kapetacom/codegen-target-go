/**
 * Copyright 2023 Kapeta Inc.
 * SPDX-License-Identifier: MIT
 */

import { format, Target } from '@kapeta/codegen-target';
import type { GeneratedAsset, SourceFile, GeneratedFile } from '@kapeta/codegen';
import Path from 'path';
import { exec } from '@kapeta/nodejs-process';
import { execSync } from 'child_process';
import { mergeDevcontainers } from './target/merge-devcontainers';
import { addTemplateHelpers } from './target/template-helpers';

export default class GoTarget extends Target {
    constructor(options: any) {
        super(options, Path.resolve(__dirname, '../'));
    }

    mergeFile(sourceFile: SourceFile, newFile: GeneratedFile, lastFile: GeneratedFile): GeneratedFile {
        if (sourceFile.filename === '.devcontainer/devcontainer.json') {
            return mergeDevcontainers(sourceFile, newFile, lastFile);
        }

        return super.mergeFile(sourceFile, newFile, lastFile);
    }

    protected _createTemplateEngine(data: any, context: any) {
        const engine = super._createTemplateEngine(data, context);

        addTemplateHelpers(engine, data, context);

        return engine;
    }

    protected _postProcessCode(filename: string, code: string) {
        if (filename.endsWith('.go')) {
            try {
                code = this.formatGoCode(code);
                return code;
            } catch (error: any) {
                console.log(code)
                throw new Error(`Error formatting go code for ${filename}: ${error.stderr || error.message}`);
            }

        }
        return format(filename, code);
    }

    private formatGoCode(inputCode: string): string {
        try {
            // Execute gofmt command synchronously with inputCode as input
            return execSync(`gofmt`, { input: inputCode, encoding: 'utf-8' });
        } catch (error: any) {
            // Handle errors
            throw new Error(`Error executing gofmt: ${error.stderr || error.message}`);
        }
    }

    generate(data: any, context: any): GeneratedFile[] {
        return super.generate(data, context);
    }

    async postprocess(targetDir: string, files: GeneratedAsset[]): Promise<void> {
        const anyFilesChanged = files.some((file) => file.filename.endsWith('.go') || file.filename === '.mod');
        if (!anyFilesChanged) {
            return;
        }
        // we should run go mod tidy and gofmt
        console.log('Running gofmt in %s', targetDir);
        const fmtchild = exec(`gofmt -w ${targetDir}`);
        await fmtchild.wait();
        console.log("done formatting");
        try {
            console.log('Running go mod tidy in %s', targetDir);
            const child = exec('go mod tidy', {
                cwd: targetDir,
            });

            await child.wait();
            console.log('done tidying');
        } catch (error: any) {
            // Handle errors
            throw new Error(`Error executing go mod tidy: ${error.stderr || error.message}`);
        }
    }
}
