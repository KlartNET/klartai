interface MarkdownProps {
	content: string;
}



export default function Markdown({ content }: MarkdownProps) {
	const parseMarkdown = (text: string) => {
		// 1. 기본 탈출 및 코드 블록 처리 (전체 텍스트를 조각내서 처리)
		const parts = text.split(/(```[\s\S]*?```|`[^`]+`)/g);

		return parts.map(part => {
			// 코드 블록인 경우
			if (part.startsWith("```")) {
				const code = part.slice(3, -3).replace(/&/g, "&amp;").replace(/</g, "&lt;");
				return `<pre><code>${code}</code></pre>`;
			}
			// 인라인 코드인 경우
			if (part.startsWith("`")) {
				const code = part.slice(1, -1).replace(/&/g, "&amp;").replace(/</g, "&lt;");
				return `<code>${code}</code>`;
			}

			// 일반 텍스트인 경우에만 마크다운 적용
			let html = part
				.replace(/&/g, "&amp;")
				.replace(/</g, "&lt;")
				.replace(/>/g, "&gt;");

			html = html.replace(/\*\*([^*]+)\*\*/g, "<strong>$1</strong>");
			html = html.replace(/_([^_]+)_/g, "<em>$1</em>");
			html = html.replace(/\*([^*]+)\*/g, "<em>$1</em>");
			html = html.replace(/^### (.*$)/gim, "<h3>$1</h3>");
			html = html.replace(/^## (.*$)/gim, "<h2>$1</h2>");
			html = html.replace(/^# (.*$)/gim, "<h1>$1</h1>");
			html = html.replace(/^\s*-\s+(.*$)/gim, "<ul><li>$1</li></ul>");
			html = html.replace(/\n/g, "<br/>");

			return html;
		}).join('').replace(/<\/ul><br \/><ul>/g, '').replace(/<\/ul><ul>/g, '');
	};

	return (
		<div
			className="markdown-body"
			dangerouslySetInnerHTML={{ __html: parseMarkdown(content) }}
		/>
	);
}