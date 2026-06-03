import type { Model, Message } from "./@types/types";

export const api = {
	async getModels(): Promise<Model[]> {
		const res = await fetch("/api/models");
		return res.json();
	},


	async chat(model: string, messages: Message[]): Promise<any> {
		const res = await fetch("/api/chat", {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({
				model,
				messages
			})
		});

		return res.json();
	}
};
