import { describe, expect, it } from 'vitest';

import type { Listener } from '../contract';
import {
	filterListeners,
	formatPortsCount,
	getListenerKey,
	groupListeners,
	summarizeListeners
} from './listeners';

const listeners: Listener[] = [
	{
		process: 'node',
		pid: 101,
		protocol: 'tcp',
		address: '127.0.0.1',
		port: 3000,
		exe: '/usr/local/bin/node'
	},
	{
		process: 'node',
		pid: 101,
		protocol: 'tcp',
		address: '0.0.0.0',
		port: 5173,
		exe: '/usr/local/bin/node'
	},
	{
		process: 'node',
		pid: 202,
		protocol: 'udp',
		address: '::1',
		port: 5353,
		exe: '/opt/node/bin/node'
	}
];

describe('groupListeners', () => {
	it('groups ports by process, PID and executable without merging distinct processes', () => {
		expect(groupListeners(listeners)).toEqual([
			{
				key: 'node\u0000101\u0000/usr/local/bin/node',
				process: 'node',
				pid: 101,
				exe: '/usr/local/bin/node',
				listeners: [listeners[0], listeners[1]]
			},
			{
				key: 'node\u0000202\u0000/opt/node/bin/node',
				process: 'node',
				pid: 202,
				exe: '/opt/node/bin/node',
				listeners: [listeners[2]]
			}
		]);
	});
});

describe('filterListeners', () => {
	it('matches every listener field case-insensitively and trims the query', () => {
		expect(filterListeners(listeners, ' NODE ')).toEqual(listeners);
		expect(filterListeners(listeners, '202')).toEqual([listeners[2]]);
		expect(filterListeners(listeners, 'UDP')).toEqual([listeners[2]]);
		expect(filterListeners(listeners, '127.0.0.1')).toEqual([listeners[0]]);
		expect(filterListeners(listeners, '5173')).toEqual([listeners[1]]);
		expect(filterListeners(listeners, '/OPT/NODE')).toEqual([listeners[2]]);
	});
});

describe('summarizeListeners', () => {
	it('counts visible services, ports and protocols', () => {
		expect(summarizeListeners(groupListeners(listeners))).toEqual({
			services: 2,
			ports: 3,
			tcp: 2,
			udp: 1
		});
	});
});

describe('getListenerKey', () => {
	it('keeps duplicate endpoint rows renderable with distinct keys', () => {
		expect(getListenerKey(listeners[0], 0)).not.toBe(getListenerKey(listeners[0], 1));
	});
});

describe('formatPortsCount', () => {
	it('uses the correct Russian plural form', () => {
		expect(formatPortsCount(1)).toBe('1 порт');
		expect(formatPortsCount(2)).toBe('2 порта');
		expect(formatPortsCount(5)).toBe('5 портов');
		expect(formatPortsCount(11)).toBe('11 портов');
		expect(formatPortsCount(21)).toBe('21 порт');
	});
});
