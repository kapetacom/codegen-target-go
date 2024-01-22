import { jwksRoutes, jwtAuthorization, JWTKeyStore } from '@kapeta/sdk-auth-jwt';
import { ConfigProvider } from '@kapeta/sdk-config';
import { Router } from 'express';

import { createProviderKeyStore } from '../../auth/keystore-provider';

export const createAuthMiddleware = async (config: ConfigProvider): Promise<Router> => {
    const providerKeyStore: JWTKeyStore = await createProviderKeyStore();

    const router = Router();
    const keyStores: JWTKeyStore[] = [providerKeyStore];

    router.use(
        jwtAuthorization({
            keyStores,
        })
    );

    router.use(jwksRoutes(providerKeyStore));

    return router;
};
