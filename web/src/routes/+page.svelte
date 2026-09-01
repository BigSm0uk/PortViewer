<script lang="ts">
	import Listeners from '../features/collector/ui/Listeners.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

{#await data.listeners}
	Loading listeners...
{:then listeners}
	<Listeners listeners={listeners} />
{/await}
