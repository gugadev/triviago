const { join, resolve } = require('path')
const webpack = require('webpack')

module.exports = {
    entry: {
        bundle: './src/main.js',
    },
    output: {
        path: join(__dirname, './public'),
        filename: 'js/[name].js',
    },
    module: {
        rules: [
            {
                test: /.js?$/,
                exclude: /node_modules/,
                use: [
                    {
                        loader: 'babel-loader',
                    }
                ],
            },
            {
                test: /.vue?$/,
                exclude: /node_modules/,
                use: [
                    {
                        loader: 'vue-loader',
                    },
                ],
            },
            {
                test: /.json?$/,
                exclude: /node_modules/,
                use: [
                    {
                        loader: 'json-loader',
                    }
                ],
            },
        ],
    }, // end module
    resolve: {
        modules: [
            resolve('/'),
            'node_modules/',
        ],
        extensions: ['.js', '.vue',],
        unsafeCache: true,
    },
    plugins: [
        new webpack.ProvidePlugin({
            Promise: 'imports-loader?this=>global!exports-loader?global.Promise!es6-promise',
        }),
        new webpack.DefinePlugin({
            'process.env.NODE_ENV': JSON.stringify('production'),
        }),
        new webpack.optimize.UglifyJsPlugin({
            compress: true,
        }),
    ],
    devtool: 'source-map',
    stats: { colors: true },
    node: {
        console: false,
        fs: 'empty',
        net: 'empty',
        tls: 'empty',
    },
}
