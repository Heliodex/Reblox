import { WindowIsMaximised } from "$lib/wailsjs/runtime"

export const prerender = true
export const ssr = false

export const load = async () => ({
	maximised: await WindowIsMaximised(),
})
