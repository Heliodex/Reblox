<script lang="ts">
	import { ContactServer, Greet } from "$lib/wailsjs/go/main/App"
	import Button from "$lib/components/Button.svelte"
	import TextField from "$lib/components/TextField.svelte"

	let name = $state("nerds")
	let text = $state("Testing")

	async function sayHi() {
		text = await Greet(name)
	}

	async function contactServer() {
		text = "Contacting server..."
		text = await ContactServer(name)
	}
</script>

<div class="p-4 grid grid-cols-2">
	<p>
		{text}
		<br />
		You are {name}
	</p>

	<div class="flex gap-2 h-min">
		<TextField
			type="text"
			bind:value={name}
			placeholder="Enter your name" />

		<Button onclick={sayHi}>Greet</Button>

		<Button onclick={contactServer}>Contact server</Button>
	</div>
</div>
