import type { Metrics, Message, Model } from "../@types/types.d.ts";

import { useState, useRef, useEffect } from 'react';

import { api } from "../api.ts";

import Markdown from "../components/Markdown.tsx";
import ModelSelector from "../components/ModelSelector/index.tsx";

import "./index.css";



export default function Home() {
	const [messages, setMessages] = useState<Message[]>([]);
	const [input, setInput] = useState('');
	const [loading, setLoading] = useState(false);
	const [metrics, setMetrics] = useState<Metrics>({ ctx: 0 });
	const [modelList, setModelList] = useState<Model[]>([]);
	const [selectedModel, setSelectedModel] = useState<string>('');
	const [isModal, setModal] = useState(false);
	const scrollRef = useRef<HTMLDivElement>(null);

	const currentModel = modelList.find(m => m.code === selectedModel);

	useEffect(() => {
		api.getModels()
			.then(data => {
				const sorted = data.sort(
					(a: Model, b: Model) =>
						b.code.localeCompare(a.code)
				);

				setModelList(sorted);

				if (sorted.length > 0)
					setSelectedModel(sorted[3].code);
			});
	}, []);

	useEffect(() => {
		scrollRef.current?.scrollTo({
			top: scrollRef.current.scrollHeight,
			behavior: "auto"
		});
	}, [messages, loading]);

	const send = async (event: React.SubmitEvent<HTMLFormElement>) => {
		event.preventDefault();
		if (!input.trim() || loading) return;

		const userMessage = {
			role: "user",
			content: input.trim()
		};
		const nextMessages = [...messages,
			userMessage
		];
		setMessages(nextMessages);

		setInput('');
		setLoading(true);

		try {
			const data = await api.chat(selectedModel, nextMessages);

			setMessages(prev =>
				[...prev, {
					role: "assistant",
					content: data.choices[0].message.content
				}]
			);
			setMetrics({
				ctx: data.usage.total_tokens
			});
		} catch {
			setMessages(prev =>
				[...prev, {
					role: "assistant",
					content: "오류가 발생했습니다."
				}]
			);
		} finally {
			setLoading(false);
		}
	};

	return (
		<div className="chat-container">
			<header className="chat-header">
				<button
					className="model-trigger"
					disabled={loading}
					onClick={() => setModal(true)}
				>
					{
						currentModel?.displayName || "모델 선택"
					}
					<span className="chevron">▼</span>
				</button>
				<span className="metrics">
					컨텍스트 사용량: {metrics.ctx} / {currentModel?.contextSize || 0}
				</span>
			</header>
			
			<main
				className="chat-messages"
				ref={scrollRef}
			>
				{
					messages.map((msg, i) => (
						<div
							key={i}
							className={`message ${msg.role}`}
						>
							<div className="bubble">
								<Markdown content={msg.content}/>
							</div>
						</div>
					))
				}
				{loading && (
					<div className="message assistant">
						<div className="bubble thinking">
							<span>●</span><span>●</span><span>●</span>
						</div>
					</div>
				)}
			</main>

			<form
				className="chat-input-form"
				onSubmit={send} 
			>
				<input
					placeholder="메시지 입력"
					value={input}
					disabled={loading}
					onChange={e => setInput(e.target.value)}
				/>
				<button disabled={loading || !input.trim()}>
					전송
				</button>
			</form>

			<ModelSelector 
				isOpen={isModal}
				models={modelList}
				selectedModel={selectedModel}
				onClose={() => setModal(false)}
				onSelect={(code) => {
					setSelectedModel(code);
					setMessages([]);
					setMetrics({ ctx: 0 });
				}}
			/>
		</div>
	);
}