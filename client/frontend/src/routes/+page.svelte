<script lang="ts">
	import { Greet } from "$lib/wailsjs/go/main/App"

	let name = $state("")
	let text = $state("Testing")

	async function sayHi() {
		text = await Greet(name)
	}
</script>

<h1>Welcome to SvelteKit</h1>
<p>
	Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a>
	to read the documentation
</p>

<p>{text}</p>

<input type="text" bind:value={name} placeholder="Enter your name" />
<button onclick={sayHi}>Greet</button>
