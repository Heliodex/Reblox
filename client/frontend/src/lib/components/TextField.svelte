<script lang="ts">
	let { children = null, value = $bindable(), ...rest } = $props()

	// when the input value changes, update the value prop
	function onInput(event: InputEvent & { target: HTMLInputElement }) {
		value = event.target.value
	}
</script>

<fluent-text-field {value} {onInput} {...rest}>
	{@render children()}
</fluent-text-field>
