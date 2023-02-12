import { useState } from 'react'

const submit = (account: string, password: string) => {
	const user = new FormData()
	user.append('account', account)
	user.append('password', password)
	fetch('http://127.0.0.1:4000/login', {
		method: 'POST',
		body: user,
	})
	.then(res => {
		if (res.ok){
			return res.json()
		}
		throw new Error('请求失败')
	})
	.then(res => {
		console.log(res.user)
	})
	.catch(err => {
		console.error(err)
	})
}

export default function Login () {
	const [account, setAccount] = useState<string>('visitant')
	const [password, setPassword] = useState<string>('1')
	return (
		<section>
			<form action=''>
				<label htmlFor='account'>account:
					<input type='text' name='account' id='account' value={ account } onChange={ (e) => setAccount(e.target.value) } />
				</label>password:
				<label htmlFor='password'>
					<input type='text' name='password' id='password' value={ password } onChange={ (e) => setPassword(e.target.value) } />
				</label>
				<button onClick={ () => submit(account, password) }>Submit</button>
			</form>
		</section>
	)
};
