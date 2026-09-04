import { afterEach, describe, expect, it, vi } from 'vitest';

import { listenerQuery } from './query';

describe('listenerQuery', () => {
	afterEach(() => {
		vi.unstubAllGlobals();
	});

	it('reports the HTTP error without parsing an empty body as JSON', async () => {
		vi.stubGlobal(
			'fetch',
			vi.fn(async () => new Response('', { status: 502, statusText: 'Bad Gateway' }))
		);

		await expect(listenerQuery()).rejects.toThrow('Failed to load listeners: 502 Bad Gateway');
	});
});
