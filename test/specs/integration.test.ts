/**
 * Copyright 2023 Kapeta Inc.
 * SPDX-License-Identifier: MIT
 */

import Path from 'path';
import {describe, test, beforeEach} from '@jest/globals';

import {
    BlockCodeGenerator,
    CodeGenerator,
    CodegenHelpers,
    CodeWriter,
    GeneratedFile,
    MODE_MERGE
} from '@kapeta/codegen';
import Target from '../../src';
import FS from "fs";

describe('blocks', () => {
    beforeEach(() => {
        global.console = require('console');
    });

    test('minimal', async () => {
        const basedir = Path.resolve(__dirname, '../resources/examples/minimal');
        const data = require('../resources/examples/minimal.kapeta.yml');

        const target = new Target({
            basePackage: 'github.com/kapeta/minimal',
        });
        return testCodeGenFor(target, new BlockCodeGenerator(data), basedir);
    });

    test('todo', async () => {
        const basedir = Path.resolve(__dirname, '../resources/examples/todo');
        const data = require('../resources/examples/todo.kapeta.yml');
        const target = new Target({
            basePackage: 'github.com/kapeta/todo',

        });
        return testCodeGenFor(target, new BlockCodeGenerator(data), basedir);
    });

    test('users', async () => {
        const basedir = Path.resolve(__dirname, '../resources/examples/users');
        const data = require('../resources/examples/users.kapeta.yml');

        const target = new Target({
            basePackage: 'github.com/kapeta/users',
        });
        return testCodeGenFor(target, new BlockCodeGenerator(data), basedir);
    });
});


export async function testCodeGenFor(target: any, generator: CodeGenerator, basedir: string) {
    const results = await generator.generateForTarget(target);
    /* eslint-disable-next-line @typescript-eslint/no-var-requires */
    const {expect} = require('@jest/globals');
    let allFiles = CodegenHelpers.walkDirectory(basedir);
    if (allFiles.length === 0 || process?.env?.FORCE_GENERATE) {
        const writer = new CodeWriter(basedir, {skipAssetsFile: true});
        console.log('No files found in directory: %s - generating output', basedir);
        writer.write(results);
        allFiles = CodegenHelpers.walkDirectory(basedir);
    }

    const mergeFiles: string[] = [];

    results.files.forEach((result: GeneratedFile) => {
        if (result.mode === MODE_MERGE) {
            mergeFiles.push(result.filename);
        }
        const fullPath = Path.join(basedir, result.filename);
        const expected = FS.readFileSync(fullPath).toString();
        const stat = FS.statSync(fullPath);
        console.log(`Comparing files: ${fullPath}`);
        expect(toUnixPermissions(stat.mode)).toBe(result.permissions);
        if (expected !== result.content) {
            if (result.filename !== 'go.mod') {
                console.log(`Expected: ${expected}`);
                console.log(`Result: ${result.content}`);
                // fail the test
                expect(expected).toBe(result.content);
            }
        }

        const ix = allFiles.indexOf(fullPath);
        expect(allFiles).toContain(fullPath);
        if (ix > -1) {
            allFiles.splice(ix, 1);
        }
    });
    // Also verify the merges have created merge cache files
    allFiles
        .filter((path: any) => path.includes('/.kapeta/merged/'))
        .forEach((path: any) => {
            const [, filename] = path.split('/.kapeta/merged/');
            if (mergeFiles.includes(filename)) {
                allFiles.splice(allFiles.indexOf(path), 1);
                return;
            }
        });
    // remove files ending with go.sum from allFiles
    allFiles = allFiles.filter((path: any) => !path.endsWith('go.sum'));
    expect(allFiles).toEqual([]);
}

function toUnixPermissions(statsMode: number) {
    return (statsMode & parseInt('777', 8)).toString(8);
}