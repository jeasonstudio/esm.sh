package server

const indexHTML = `<!DOCTYPE html>
<html>
<head>
    <meta charSet="utf-8" />
    <meta name="viewport" content="user-scalable=no,initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0" />
    <title>ESM</title>
    <meta name="description" content="A fast, global content delivery network for ES Modules. All modules are transformed to ESM by esbuild in NPM." />
    <meta name="keywords" content="esm,npm,deno,global,cdn,package,esbuild,transform" />
    <style>
        // https://esm.sh/highlightjs@9.16.2/styles/github.css
        .hljs{display:block;overflow-x:auto;padding:.5em;color:#333;background:#f8f8f8}.hljs-comment,.hljs-quote{color:#998;font-style:italic}.hljs-keyword,.hljs-selector-tag,.hljs-subst{color:#333;font-weight:700}.hljs-literal,.hljs-number,.hljs-tag .hljs-attr,.hljs-template-variable,.hljs-variable{color:teal}.hljs-doctag,.hljs-string{color:#d63369}.hljs-section,.hljs-selector-id,.hljs-title{color:#d63369;font-weight:700}.hljs-subst{font-weight:400}.hljs-class .hljs-title,.hljs-type{color:#458;font-weight:700}.hljs-attribute,.hljs-name,.hljs-tag{color:teal;font-weight:400}.hljs-link,.hljs-regexp{color:#009926}.hljs-bullet,.hljs-symbol{color:#d63369}.hljs-built_in,.hljs-builtin-name{color:teal}.hljs-meta{color:#999;font-weight:700}.hljs-deletion{background:#fdd}.hljs-addition{background:#dfd}.hljs-emphasis{font-style:italic}.hljs-strong{font-weight:700}

        /* esm.sh */
        * {
            margin: 0;
            padding: 0;
            border: none;
            outline: none;
            font: inherit;
            font-size: 100%%;
            vertical-align: baseline;
            background: transparent;
        }

        html {
            font-size: 15px;
        }

        body {
            width: 90%%;
            max-width: 840px;
            margin: 0 auto;
            padding: 120px 0 60px;
            font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Helvetica Neue', Helvetica, Roboto, Ubuntu, Tahoma, Arial, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'PingFang SC', 'Hiragino Sans GB', 'Heiti SC', 'Microsoft YaHei', 'WenQuanYi Micro Hei', sans-serif;
            color: #333;
            text-rendering: optimizeLegibility;
        }

        ul,
        ol {
            list-style: none;
        }

        a {
            color: #000;
            text-decoration: none;
        }

        img {
            max-width: 100%%;
            height: auto;
            border-radius: 6px;
        }

        strong, b {
            font-weight: 600;
        }

        em, i {
            font-style: italic;
        }

        h1, h2, h3, h4, h5 {
            font-weight: 500;
            line-height: 1.27;
        }

        h1 {
            padding-bottom: 9px;
            border-bottom: 1px solid #eee;
        }

        h1 {
            margin-top: 0.6rem;
            font-size: 2.4rem;
            font-weight: 800;
        }

        h2 {
            margin-top: 2.4rem;
            font-size: 1.8rem;
        }

        h3 {
            margin-top: 2rem;
            font-size: 1.5rem;
        }

        h4 {
            font-size: 1.2rem;
        }

        h5 {
            font-size: 1rem;
        }

        p {
            margin-top: 1.27rem;
            line-height: 1.5;
        }

        a {
            color: #d63369;
            box-shadow: 0 1px 0 0 currentColor;
        }
        a:hover {
            box-shadow: none;
        }

        code {
            font-family: 'Dank Mono', 'Source Code Pro', 'Courier Prime Code', 'Liberation Mono', Consolas, Menlo, monospace;
        }

        pre {
            box-sizing: border-box;
            overflow-x: auto;
            width: 100%%;
            margin-top: 1.27rem;
            border-radius: 6px;
            line-height: 1.5;
            color: #333;
            background-color: #f8f8f8;
            white-space: pre;
            -webkit-overflow-scrolling: touch;
        }
        pre > code {
            display: block;
            padding: 1.2rem!important;
        }

        :not(pre) > code {
            display: inline;
            white-space: pre-wrap;
            color: #d63369;
        }
        :not(pre) > code::before,
        :not(pre) > code::after {
            color: currentColor;
            content: '%s'
        }
        pre > code .bash_prompt {
            color: #bbb;
            user-select: none;
        }

        details {
            margin: 1.5rem 0;
            padding: 0.5rem 1rem;
            background: #fafafa;
            border: 1px solid #eaeaea;
            border-radius: 6px;
        }

        details[open] {
            overflow: hidden;
        }

        details > summary {
            font-weight: 500;
            outline: none;
            cursor: pointer;
        }

        blockquote {
            color: #666666;
            background: #fafafa;
            border: 1px solid #eaeaea;
            border-radius: 6px;
            padding: 0 1.25rem;
            margin: 1.5rem 0;
        }

        hr {
            border: 0;
            border-top: 1px solid #eaeaea;
            margin: 1.25rem 0;
        }

        ul, ol {
            padding-left: 1.5rem;
            margin-top: 1.27rem;
        }

        ol {
            list-style-type: decimal;
        }

        li {
            margin-top: 0.6rem;
        }

        ul li:before {
            position: absolute;
            margin-left: -1rem;
            color: #aaa;
            content: '-';
        }

        h1 {
            display: flex;
            align-items: center;
        }

        h1 a {
            position: relative;
            display: inline-block;
            width: 28px;
            height: 28px;
            margin-left: 15px;
            color: #555;
            box-shadow: none;
        }

        h1 a svg {
            position: absolute;
            display: inline-block;
            top: 0;
            left: 0;
            width: 28px;
            height: 28px;
        }

        h1 a:hover {
           color: #000;
        }

        .test {
            display: none;
        }

        .test p {
            display: inline-block;
        }
    </style>
</head>
<body>
    <h1>
        <strong>ESM</strong>
        <a href="https://github.com/postui/esm.sh">
            <svg fill="currentColor" viewBox="0 0 24 24">
                <title>Github</title>
                <path fill-rule="evenodd" d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z" clip-rule="evenodd"></path>
            </svg>
        </a>
    </h1>
    <main>
        <p><em style="color: #999;">Loading...</em></p>
    </main>
    <div class="test">
        <h2>Test</h2>
        <p><strong>React</strong>: <span id="react">❌</span></p>
        <p><strong>React(bundle)</strong>: <span id="reactb">❌</span></p>
        <p><strong>Preact</strong>: <span id="preact"><span className="x">❌</span></span></p>
        <p><strong>Vue2</strong>: <span id="vue">❌</span></p>
        <p><strong>Vue3</strong>: <span id="vue3">❌</span></p>
    </div>
    <script type="module">
        import marked from '/marked';
        import hljs from '/highlight.js/lib/core';
        import javascript from '/highlight.js/lib/languages/javascript';
        import json from '/highlight.js/lib/languages/json';
        import bash from '/highlight.js/lib/languages/bash';

        const mainEl = document.querySelector('main');
        mainEl.innerHTML = marked.parse(%s);
        mainEl.removeChild(mainEl.querySelector('h1'));
        hljs.registerLanguage('javascript', javascript);
        hljs.registerLanguage('json', json);
        hljs.registerLanguage('bash', hljs => {
            const l = bash(hljs)
            l.keywords.built_in = 'cd git sh esm deno aleph'
            return l
        });
        hljs.initHighlighting();
        document.querySelectorAll('code.language-bash').forEach(block => {
            for (let i = 0; i < block.childNodes.length; i++) {
                const child = block.childNodes[i]
                if (child.nodeName === '#text') {
                    const text = child.textContent
                    if (text == '$ ') {
                        block.insertBefore(bashPromptSpan(), child)
                        block.removeChild(child)
                    } else {
                        const texts = text.split('\n$ ')
                        const n = texts.length
                        if (n > 1) {
                            for (let j = 0; j < n; j++) {
                                const t = texts[j]
                                if (t) {
                                    const node = document.createTextNode(t + '\n')
                                    block.insertBefore(node, child)
                                } else if (j == 0) {
                                    const node = document.createTextNode('\n')
                                    block.insertBefore(node, child)
                                }
                                if (j > 0) {
                                    block.insertBefore(bashPromptSpan(), child)
                                }
                            }
                            block.removeChild(child)
                        }
                    }
                }
            }
        });

        function bashPromptSpan(prompt = '$') {
            const span = document.createElement('span')
            span.className = 'bash_prompt'
            span.innerText = prompt + ' '
            return span
        }

        mainEl.querySelectorAll('img').forEach(img => {
            const src = img.getAttribute('src')
            if (src.startsWith('./assets/')) {
                img.src = 'https://denopkg.com/postui/esm.sh/' + src.replace('./', '')
            }
        })

        if (location.hostname === 'localhost') {
            document.querySelector('.test').style.display = "block";

            (async () => {
                const { createElement } = await import('/react?dev');
                const { render } = await import('/react-dom?dev');
                render(createElement('span', null, '✅'), document.querySelector('#react'));
            })();

            (async () => {
                const { createElement } = await import('/[react,react-dom]/react?dev');
                const { render } = await import('/[react,react-dom]/react-dom?dev');

                render(createElement('span', null, '✅'), document.querySelector('#reactb'));
            })();

            (async () => {
                const { h, render } = await import('/preact?dev');
                const { useEffect } = await import('/preact/hooks?dev');

                const el = document.querySelector('#preact');
                function App() {
                    useEffect(() => {
                        el.removeChild(el.querySelector('.x'))
                    }, [])
                    return h('span', null, '✅')
                }
                render(h('span', null, '✅'), el);
            })();

            (async () => {
                const {default: Vue} = await import('/vue?dev');

                new Vue({
                    el: '#vue',
                    render(h) {
                        return h('span', null, '✅')
                    }
                });
            })();

            (async () => {
                const { createApp, h } = await import('/vue@3.0.0?dev');

                createApp({
                    render() {
                        return h('span', {}, '✅')
                    }
                }).mount('#vue3');
            })();
        }
    </script>
    <script nomodule>
        const mainEl = document.querySelector('main');
        mainEl.innerHTML = '<p><em style="color: #999;">nomodule, please upgrade your browser...</em></p>'
    </script>
</body>
</html>
`
