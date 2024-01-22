//
// GENERATED SOURCE - DO NOT EDIT
//
import {
    PostgresDB as $PostgresDB,
    createPostgresDBClient as $createPostgresDBClient,
} from '@kapeta/sdk-sqldb-postgresql';
import { Prisma, PrismaClient } from './clients/users';
import { ConfigProvider } from '@kapeta/sdk-config';

export type UsersDBClient = PrismaClient;

function createClient(opts: Prisma.PrismaClientOptions): UsersDBClient {
    return new PrismaClient(opts);
}

/*
 * Create a ready and connected UsersDB client.
 *
 * See https://github.com/kapetacom/sdk-nodejs-sqldb-postgresql for more information.
 */
export const createUsersDBClient = (config: ConfigProvider) => {
    return $createPostgresDBClient<UsersDBClient>(config, 'users', createClient);
};

/**
 * Use this class to access the users database.
 * Also see {@link createUsersDBClient} as the recommended way of connecting to the database.
 *
 * See https://github.com/kapetacom/sdk-nodejs-sqldb-postgresql for more information.
 */
export class UsersDB extends $PostgresDB<UsersDBClient> {
    constructor() {
        super('users');
    }

    createClient(opts: any): UsersDBClient {
        return createClient(opts);
    }
}
