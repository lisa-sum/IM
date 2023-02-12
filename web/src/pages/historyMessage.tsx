const getHistoryMessage = () => {

}

export default function HistoryMessage ({ historyMessage }: { historyMessage: MessageBasic[] }) {
	return (
		<section>
			<button onClick={getHistoryMessage}>getHistoryMessage</button>
			<ol>
				{
					historyMessage.map((item:MessageBasic, index:number) => {
						return <li key={ index }>
							<span>用户ID:{ item.user_identity }</span>
							<p>房间ID:{ item.room_identity }</p>
							<p>消息:{ item.data }</p>
							<p>创建时间:{ item.created_at }</p>
							<p>更新时间:{ item.updated_at }</p>
						</li>
					})
				}
			</ol>
		</section>
	)
};
