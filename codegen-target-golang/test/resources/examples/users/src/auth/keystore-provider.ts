import { ensureFileKeystore, JWTKeyStoreInternal } from '@kapeta/sdk-auth-jwt';

/**
 * Create a keystore provider for the JWT provider.
 *
 * This is used to sign and verify JWT tokens.
 *
 * You should adjust this implementation to fit your needs.
 */
export const createProviderKeyStore = async () => {
    const jwtConfig = {
        issuer: 'https://example.auth.kapeta.com',
        audience: 'https://example.kapeta.com',
        file: 'jwks.json',
    };

    const jwksStore = await ensureFileKeystore(jwtConfig.file);
    return new JWTKeyStoreInternal(jwksStore, jwtConfig.issuer, jwtConfig.audience);
};
