export interface Metrics {
	ctx: number;
}

export interface Message {
	role: string;
	content: string;
}

export interface Model {
	code: string;
	displayName: string;
	contextSize: string;
	source: string;
	parameters: string;
	tags?: string[];
}