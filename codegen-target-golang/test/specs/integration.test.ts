import Path from 'path';
import { describe, test, beforeEach } from '@jest/globals';

import { CodegenHelpers, BlockCodeGenerator } from '@kapeta/codegen';
import Target from '../../src';

describe('blocks', () => {
    let target: Target;

    beforeEach(() => {
        target = new Target({
            basePackage: 'github.com/kapeta/todo',
        });
    });

    test('todo', async () => {
        const basedir = Path.resolve(__dirname, '../resources/examples/todo');
        const data = require('../resources/examples/todo.kapeta.yml');

        return CodegenHelpers.testCodeGenFor(target, new BlockCodeGenerator(data), basedir);
    });

    test('users', async () => {
        const basedir = Path.resolve(__dirname, '../resources/examples/users');
        const data = require('../resources/examples/users.kapeta.yml');

        return CodegenHelpers.testCodeGenFor(target, new BlockCodeGenerator(data), basedir);
    });
});
