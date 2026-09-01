import { collectorStore } from '../features/collector/collector.svelte';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	return {
		listeners: collectorStore.fetchCollector(fetch)
	};
};
