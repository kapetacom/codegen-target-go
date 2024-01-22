import { createExternalKeyStore, jwtAuthorization, JWTKeyStore } from '@kapeta/sdk-auth-jwt';
import { ConfigProvider } from '@kapeta/sdk-config';
import { Router } from 'express';

export const createAuthMiddleware = async (config: ConfigProvider): Promise<Router> => {
    const router = Router();
    const keyStores: JWTKeyStore[] = [await createExternalKeyStore('auth', config)];

    router.use(
        jwtAuthorization({
            keyStores,
        })
    );

    return router;
};
