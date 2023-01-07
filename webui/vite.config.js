import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig(({command, mode, ssrBuild}) => {
	const ret = {
		plugins: [vue({
			template: {
			  compilerOptions: {
				isCustomElement: (tag) => ['main-footer' ].includes(tag),
			  }
			}
		  })],
		resolve: {
			alias: {
				'@': fileURLToPath(new URL('./src', import.meta.url))
			}
		},
	};
	ret.define = {
		"__API_URL__": JSON.stringify("http://localhost:3000"),
	};
	return ret;
})