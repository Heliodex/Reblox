<script lang="ts">
	import {
		WindowMinimise,
		WindowToggleMaximise,
		WindowIsMaximised,
		WindowReload,
		Quit
	} from "$lib/wailsjs/runtime"
	import "uno.css"
	import "../global.styl"

	const { children, data } = $props()

	let windowMaximised = $state(data.maximised)
</script>

<header
	class="bg-neutral-8 h-8 grid grid-cols-2 items-center select-none"
	role="button"
	tabindex="-1"
	ondblclick={WindowToggleMaximise}>
	<div class="h-full flex justify-left">
		<a href="/" class="px-4">
			<img src="/reblox.svg" alt="Reblox Logo" class="h-6.5 pt-1.75" />
		</a>
		<button onclick={WindowReload} class="min-w-9!">
			<img src="/refresh.svg" alt="refresh" class="pt-0.5" />
		</button>
	</div>

	<div class="h-full flex justify-right">
		<button onclick={WindowMinimise}>
			<img src="/minimise.svg" alt="Minimise" />
		</button>
		<button
			onclick={async () => {
				WindowToggleMaximise()
				await new Promise(resolve => setTimeout(resolve, 100)) // jank
				windowMaximised = await WindowIsMaximised()
			}}>
			{#if windowMaximised}
				<img src="/restore.svg" alt="Restore down" />
			{:else}
				<img src="/maximise.svg" alt="Maximise" />
			{/if}
		</button>
		<button onclick={Quit} class="hover:bg-red-7! active:bg-red-6!">
			<img src="/quit.svg" alt="Quit" />
		</button>
	</div>
</header>

<main>
	{@render children()}
</main>

<style lang="stylus">
	header
		--wails-draggable drag

	button
	a
		--wails-draggable no-drag
		@apply h-full min-w-12 bg-transparent transition
		cursor pointer
		color white
		text-decoration none
		&:hover
			@apply bg-neutral-7
		&:active
			@apply bg-neutral-6
</style>
