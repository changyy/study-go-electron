const HtmlWebpackPlugin = require('html-webpack-plugin')
const path = require('path');
const webpack = require('webpack');
const FileManagerPlugin = require('filemanager-webpack-plugin');

module.exports = {
    entry: path.join(__dirname, 'src/js/index.js'),
    output: {
        path: path.join(__dirname, 'dist'),
        filename: 'bundle.js'
    },
    plugins: [
        // https://webpack.js.org/plugins/hot-module-replacement-plugin/
        //new webpack.HotModuleReplacementPlugin(),
        // https://github.com/jantimon/html-webpack-plugin
        new HtmlWebpackPlugin({
            title: 'My Go Electron Project',
            filename: 'index.html',
            template: 'src/index.html',
        }),
        new FileManagerPlugin({
            events: {
                onEnd: {
                    delete: [
                        {
                            source: path.join(__dirname, '../resources/app'),
                            options: {
                                force: true,
                            },
                        }
                    ],
                    mkdir: [
                        path.join(__dirname, '../resources/app')
                    ],
                    copy: [
                        {
                            source: path.join(__dirname, 'dist'),
                            destination: path.join(__dirname, '../resources/app'), 
                        },
                    ]
                }
            }
        }),
    ]
}
