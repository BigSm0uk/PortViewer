import type { Listener } from '../contract';

export interface ListenerGroup {
	key: string;
	process: string;
	pid: number;
	exe: string;
	listeners: Listener[];
}

export interface ListenerSummary {
	services: number;
	ports: number;
	tcp: number;
	udp: number;
}

export function formatPortsCount(count: number): string {
	const absoluteCount = Math.abs(count);
	const lastTwoDigits = absoluteCount % 100;
	const lastDigit = absoluteCount % 10;
	let label = 'портов';

	if (lastDigit === 1 && lastTwoDigits !== 11) {
		label = 'порт';
	} else if (lastDigit >= 2 && lastDigit <= 4 && (lastTwoDigits < 12 || lastTwoDigits > 14)) {
		label = 'порта';
	}

	return `${count} ${label}`;
}

export function getListenerKey(listener: Listener, index: number): string {
	return `${listener.protocol}\u0000${listener.address}\u0000${listener.port}\u0000${index}`;
}

export function summarizeListeners(groups: ListenerGroup[]): ListenerSummary {
	const summary: ListenerSummary = {
		services: groups.length,
		ports: 0,
		tcp: 0,
		udp: 0
	};

	for (const group of groups) {
		for (const listener of group.listeners) {
			summary.ports += 1;
			const protocol = listener.protocol.toLocaleLowerCase();
			if (protocol === 'tcp') summary.tcp += 1;
			if (protocol === 'udp') summary.udp += 1;
		}
	}

	return summary;
}

export function filterListeners(listeners: Listener[], query: string): Listener[] {
	const normalizedQuery = query.trim().toLocaleLowerCase();

	if (!normalizedQuery) return listeners;

	return listeners.filter((listener) =>
		[
			listener.process,
			listener.pid,
			listener.protocol,
			listener.address,
			listener.port,
			listener.exe
		].some((value) => String(value).toLocaleLowerCase().includes(normalizedQuery))
	);
}

export function groupListeners(listeners: Listener[]): ListenerGroup[] {
	const groups = new Map<string, ListenerGroup>();

	for (const listener of listeners) {
		const key = `${listener.process}\u0000${listener.pid}\u0000${listener.exe}`;
		const group = groups.get(key);

		if (group) {
			group.listeners.push(listener);
			continue;
		}

		groups.set(key, {
			key,
			process: listener.process,
			pid: listener.pid,
			exe: listener.exe,
			listeners: [listener]
		});
	}

	return [...groups.values()];
}
