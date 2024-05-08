<script lang="ts">
	import { ContactServer, Greet } from "$lib/wailsjs/go/main/App"

	let name = $state("")
	let text = $state("Testing")

	async function sayHi() {
		text = await Greet(name)
	}

	async function contactServer() {
		text = "Contacting server..."
		text = await ContactServer("testing")
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

<button onclick={contactServer}>Contact server</button>
