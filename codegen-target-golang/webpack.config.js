const Path = require('path');
const FS = require('fs');
const YAML = require('yaml');

const blockInfo = YAML.parse(FS.readFileSync(Path.resolve(__dirname, './kapeta.yml')).toString());
const packageJson = require('./package.json');

module.exports = {
    entry: {
        [`${blockInfo.metadata.name}:${packageJson.version}`]: {
            import: Path.resolve(__dirname, './src/web'),
            filename: `${blockInfo.metadata.name}.js`,
        },
    },
    output: {
        path: Path.join(process.cwd(), 'web'),
        filename: '[name].js',
        library: {
            name: `Kapeta.languageTargets["[name]"]`,
            type: 'assign',
            export: 'default',
        },
    },
    module: {
        rules: [
            {
                test: /\.(ts|tsx)$/,
                loader: 'babel-loader',
                options: {
                    presets: ['@babel/env', '@babel/typescript', '@babel/react'],
                    plugins: [
                        ['@babel/plugin-proposal-decorators', { legacy: true }],
                        ['@babel/plugin-proposal-private-methods', { loose: true }],
                        ['@babel/plugin-proposal-private-property-in-object', { loose: true }],
                        ['@babel/plugin-proposal-class-properties', { loose: true }],
                        '@babel/proposal-object-rest-spread',
                    ],
                },
            },
            {
                test: /\.less$/,
                use: ['style-loader', 'css-loader', 'less-loader'],
                include: Path.resolve(__dirname, './src'),
            },
            {
                test: /\.ya?ml$/,
                use: ['yaml-loader'],
                include: Path.resolve(__dirname, './'),
            },
        ],
    },
    devtool: process.env.NODE_ENV === 'production' ? 'source-map' : 'inline-source-map',
    resolve: {
        extensions: ['.js', '.ts', '.tsx', '.less', '.yml', '.yaml'],
    },
    externals: {
        react: 'React',
        '@kapeta/ui-web-components': 'Kapeta.Components',
        '@kapeta/ui-web-types': 'Kapeta.Types',
    },
};
