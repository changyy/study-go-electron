import hehe_old_style from './helper-old-style.js';
import EntryPoint from './helper.js';

// https://stackoverflow.com/questions/34357489/calling-webpacked-code-from-outside-html-script-tag
// https://stackoverflow.com/questions/69069803/one-of-the-dist-file-is-always-empty-after-running-webpack
// https://developer.mozilla.org/zh-TW/docs/Web/JavaScript/Reference/Statements/export

//window.hehe = hehe_old_style
window.hehe = EntryPoint.hehe
