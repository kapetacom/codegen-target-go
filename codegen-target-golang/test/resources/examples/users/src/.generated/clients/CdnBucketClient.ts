//
// GENERATED SOURCE - DO NOT EDIT
//

import type { Readable } from 'stream';
import * as minio from 'minio';
import { ConfigProvider, ResourceInfo } from '@kapeta/sdk-config';

export const RESOURCE_TYPE = 'kapeta/resource-type-cloud-bucket';
export const PORT_TYPE = 'http';
export const RESOURCE_NAME = 'cdnBucket';
export const DEFAULT_BUCKET = 'cdn-bucket';

export async function createCdnBucketClient(config: ConfigProvider): Promise<CdnBucketClient> {
    const resourceInfo = await config.getResourceInfo(RESOURCE_TYPE, PORT_TYPE, RESOURCE_NAME);
    if (!resourceInfo) {
        throw new Error(`Could not find resource info for ${RESOURCE_TYPE}#${PORT_TYPE} ${RESOURCE_NAME}`);
    }
    const client = new CdnBucketClient(resourceInfo);
    // Check that bucket exists as a way to validate credentials
    try {
        await client.bucketExists();
    } catch (err) {
        throw new AggregateError([new Error(`Could not connect to bucket resource ${RESOURCE_NAME}`), err]);
    }
    return client;
}

export class CdnBucketClient {
    /**
     * The underlying minio client. Use this if you need to access the minio client directly.
     * Remember to pass the client.bucketName as the first argument to all methods.
     */
    public minio: minio.Client;
    public bucketName: string = DEFAULT_BUCKET;

    constructor(config: ResourceInfo) {
        if (!config.credentials) {
            throw new Error(`Missing credentials for ${config.type}#${config.port}`);
        }

        const fullName = config.options?.fullName;
        if (fullName) {
            this.bucketName = fullName.replaceAll(/[^a-z0-9-]+$/g, '-');
            if (this.bucketName !== fullName) {
                console.warn(`Bucket name ${fullName} contains invalid characters. Using ${this.bucketName} instead.`);
            }
        }

        this.minio = new minio.Client({
            endPoint: config.host,
            port: parseInt(`${config.port}`, 10),
            useSSL: false,
            accessKey: config.credentials.username,
            secretKey: config.credentials.password,
        });
    }

    public async listObjects(): Promise<minio.BucketItem[]> {
        const objects = await this.minio.listObjectsV2(this.bucketName, '', true);
        return objects.toArray();
    }

    public async getObject(objectPath: string): Promise<Readable> {
        return this.minio.getObject(this.bucketName, objectPath);
    }

    public async putObject(objectPath: string, data: Readable | Buffer | string): Promise<minio.UploadedObjectInfo> {
        return this.minio.putObject(this.bucketName, objectPath, data);
    }

    public async bucketExists(): Promise<boolean> {
        return this.minio.bucketExists(this.bucketName);
    }
}
