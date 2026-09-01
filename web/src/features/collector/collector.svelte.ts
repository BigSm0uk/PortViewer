import type { Listener } from './contract';
import { listenerQuery } from './query';

export class Collector {
	private listeners = $state<Listener[]>();

	async fetchCollector(fetcher: typeof fetch = fetch): Promise<Listener[]> {
		const res = await listenerQuery(fetcher);
		this.listeners = res;
		return res;
	}
}
export const collectorStore = new Collector();
