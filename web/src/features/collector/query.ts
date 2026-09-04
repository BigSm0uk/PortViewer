import type { Listener } from './contract';

export async function listenerQuery(fetcher: typeof fetch = fetch): Promise<Listener[]> {
	const response = await fetcher('/api/listeners');

	if (!response.ok) {
		throw new Error(`Failed to load listeners: ${response.status} ${response.statusText}`);
	}

	return (await response.json()) as Listener[];
}
