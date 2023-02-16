const getHistoryMessage = () => {

}

export default function HistoryMessage ({ historyMessage }: { historyMessage: MessageBasic[] }) {
	return (
		<section>
			<button onClick={ getHistoryMessage }>getHistoryMessage</button>
			<ol>
				{
					historyMessage.map((item: MessageBasic, index: number) => {
						return <li key={ index }>
							<span>用户ID:{ item.userIdentity }</span>
							<p>房间ID:{ item.roomIdentity }</p>
							<p>消息:{ item.data }</p>
							<p>创建时间:{ item.createdAt }</p>
							<p>更新时间:{ item.updatedAt }</p>
						</li>
					})
				}
			</ol>
		</section>
	)
};
