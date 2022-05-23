import React from 'react';
import Lowlight from 'react-lowlight'
import 'highlight.js/styles/atom-one-light.css'
import styles from './index.module.sass';
import { default as jsonBeautify } from "json-beautify";


import xml from 'highlight.js/lib/languages/xml'
import json from 'highlight.js/lib/languages/json'
import javascript from 'highlight.js/lib/languages/javascript'
import yaml from 'highlight.js/lib/languages/yaml'


Lowlight.registerLanguage('xml', xml);
Lowlight.registerLanguage('json', json);
Lowlight.registerLanguage('yaml', yaml);
Lowlight.registerLanguage('javascript', javascript);
const MAXIMUM_BYTES_TO_FORMAT = 1000000; // The maximum of chars to highlight in body, in case the response can be megabytes

interface Props {
    code: string;
    showLineNumbers?: boolean;
    language?: string;
    encoding?: string
    className?: string
}

export const SyntaxHighlighter: React.FC<Props> = ({
    code,
    showLineNumbers = false,
    language = "",
    encoding,
    className = ""
}) => {
    const isBase64Encoding = encoding === 'base64';
    const formatTextBody = (body: any): string => {


        const jsonLikeFormats = ['json', 'yaml', 'yml'];
        const chunk = body.slice(0, MAXIMUM_BYTES_TO_FORMAT);
        const bodyBuf = isBase64Encoding ? atob(chunk) : chunk;

        try {
            if (jsonLikeFormats.some(format => language?.indexOf(format) > -1)) {

                return jsonBeautify(JSON.parse(bodyBuf), null, 2, 80);
            }
        } catch (error) {
            console.log(error)
        }
        return bodyBuf;
    }


    const markers = showLineNumbers ? code.split("\n").map((item, i) => {
        return {
            line: i + 1,
            className: styles.hljsMarkerLine
        }
    }) : [];

    return <div style={{ fontSize: ".75rem" }} className={styles.highlighterContainer + ` ${className}`}><Lowlight language={language ? language : ""} value={formatTextBody(code)} markers={markers} /></div>;
};

export default SyntaxHighlighter;
