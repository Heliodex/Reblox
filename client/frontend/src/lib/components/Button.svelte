<script lang="ts">
	const { children = null, ...rest } = $props()
</script>

<fluent-button {...rest}>
	{@render children()}
</fluent-button>
