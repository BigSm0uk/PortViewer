<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="flex flex-1 flex-col gap-4 p-4 pt-0">
  <div class="min-h-screen flex-1 rounded-xl bg-muted/50 md:min-h-min">{@render children()}</div>
</div>
