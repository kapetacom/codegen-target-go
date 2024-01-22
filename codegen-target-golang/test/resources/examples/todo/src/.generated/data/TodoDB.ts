//
// GENERATED SOURCE - DO NOT EDIT
//

import { MongoDB as $MongoDB, createMongoDBClient as $createMongoDBClient } from '@kapeta/sdk-nosql-mongodb';
import { Prisma, PrismaClient } from './clients/todo';
import { ConfigProvider } from '@kapeta/sdk-config';

export type TodoDBClient = PrismaClient;

function createClient(opts: Prisma.PrismaClientOptions): TodoDBClient {
    return new PrismaClient(opts);
}

/*
 * Create a ready and connected TodoDB client.
 *
 * See https://github.com/kapetacom/sdk-nodejs-nosql-mongodb for more information.
 */
export const createTodoDBClient = (config: ConfigProvider) => {
    return $createMongoDBClient<TodoDBClient>(config, 'todo', createClient);
};

/**
 * Use this class to access the todo database.
 * Also see {@link createTodoDBClient} as the recommended way of connecting to the database.
 *
 * See https://github.com/kapetacom/sdk-nodejs-nosql-mongodb for more information.
 */
export class TodoDB extends $MongoDB<TodoDBClient> {
    constructor() {
        super('todo');
    }

    createClient(opts: any): TodoDBClient {
        return createClient(opts);
    }
}
