// https://webpack.js.org/guides/production/
const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');
const path = require('path');

module.exports = merge(common, {
    mode: 'production',
    plugins: [
    ]
});
