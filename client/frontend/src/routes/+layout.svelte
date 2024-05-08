<script lang="ts">
	import {
		WindowMinimise,
		WindowToggleMaximise,
		WindowIsMaximised,
		WindowReloadApp,
		Quit
	} from "$lib/wailsjs/runtime"
	import "$lib/fluent"
	import "uno.css"
	import "../global.styl"

	const { children, data } = $props()

	let windowMaximised = $state(data.maximised)
	let focussed = $state(true)

	window.addEventListener("keydown", e => {
		if (e.key === "F5") WindowReloadApp() // it doesn't happen automatically
	})
</script>

<svelte:window
	onresize={async () => {
		windowMaximised = await WindowIsMaximised()
	}}
	onfocus={() => {
		focussed = true
	}}
	onblur={() => {
		focussed = false
	}} />

<header
	class="h-8 grid grid-cols-2 items-center select-none transition-300 {focussed ? "bg-neutral-9" : "bg-neutral-9/50"}"
	style="--focus: {focussed ? "1" : "0.6"}"
	role="button"
	tabindex="-1"
	ondblclick={WindowToggleMaximise}>
	<div class="h-full flex justify-left">
		<a href="/" class="px-4">
			<img src="/reblox.svg" alt="Reblox Logo" class="h-6.5 pt-1.75" />
		</a>
		<button onclick={WindowReloadApp}>
			<img src="/refresh.svg" alt="Refresh" class="pt-0.5 focus" />
		</button>
		{#if history.length > 1}
			<button onclick={history.back}>
				<img src="/arrow.svg" alt="Back" class="pb-0.5 focus" />
			</button>
		{/if}
		{#if history.length > 1}
			<button onclick={history.forward}>
				<img
					src="/arrow.svg"
					alt="Forward"
					class="pt-0.5 transform rotate-180 focus" />
			</button>
		{/if}
	</div>

	<div class="h-full flex justify-right">
		<button onclick={WindowMinimise}>
			<img src="/minimise.svg" alt="Minimise" class="focus"/>
		</button>
		<button
			onclick={async () => {
				WindowToggleMaximise()
			}}>
			{#if windowMaximised}
				<img src="/restore.svg" alt="Restore down" class="focus"/>
			{:else}
				<img src="/maximise.svg" alt="Maximise" class="focus"/>
			{/if}
		</button>
		<button onclick={Quit} class="hover:bg-red-7! active:bg-red-6!">
			<img src="/quit.svg" alt="Quit" class="focus" />
		</button>
	</div>
</header>

<main class:rounded-2={!windowMaximised}>
	{@render children()}
</main>

<style lang="stylus">
	header
		--wails-draggable drag

	.focus
		@apply transition-300
		opacity var(--focus)

	button
	a
		--wails-draggable no-drag
		@apply h-full min-w-12 bg-transparent transition
		border none
		cursor pointer
		color white
		text-decoration none
		&:hover
			@apply bg-neutral-8
		&:active
			@apply bg-neutral-7
</style>
