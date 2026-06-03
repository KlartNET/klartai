import type { Model } from "../../@types/types.d.ts";

import "./index.css";



interface ModelSelectorProps {
	isOpen: boolean;
	onClose: () => void;
	models: Model[];
	selectedModel: string;
	onSelect: (code: string) => void;
}

export default function ModelSelector({ isOpen, onClose, models, selectedModel, onSelect }: ModelSelectorProps) {
	if (!isOpen) return null;

	return (
		<div className="modal-overlay" onClick={onClose}>
			<div className="model-modal" onClick={e => e.stopPropagation()}>
				<div className="modal-header">
					<div className="title-group">
						<h3>모델 선택</h3>
					</div>
					<button className="close-btn" onClick={onClose}>&times;</button>
				</div>
				<div className="model-list">
					{models.map(m => (
						<div
							key={m.code}
							className={`model-card ${selectedModel === m.code? 'selected' : ''}`}
							onClick={() => { onSelect(m.code); onClose(); }}
						>
							<div className="model-main-info">
								<div className="name-row">
									<span className="model-name">
										{ m.displayName }
									</span>
									{m.source &&
										<span className="model-source">
											{ m.source }
										</span>
									}
								</div>
								<div className="meta-row">
									{m.parameters &&
										<span className="meta-item">
											파라미터: { m.parameters }
										</span>
									}
									<span className="meta-item">
										최대 컨텍스트: {parseInt(m.contextSize).toLocaleString()}
									</span>
								</div>
								<div className="tag-row">
									{
										m.tags?.map(
											t => <span key={t} className="tag-pill">{t}</span>
										)
									}
								</div>
							</div>
						</div>
					))}
				</div>
			</div>
		</div>
	);
}
