import { useState } from 'react'
import { useAppDispatch } from '@/utils/hooks/index'
import { updateUserProfile } from '@/features/profile'
import { useNavigate } from 'react-router-dom'

export default function Login () {
	const [account, setAccount] = useState<string>('student')
	const [password, setPassword] = useState<string>('123456')
	const dispatch = useAppDispatch()
	const navigate = useNavigate()
	const submit = (account: string, password: string) => {
		const user = new FormData()
		user.append('account', account)
		user.append('password', password)
		fetch('http://127.0.0.1:4000/login', {
			method: 'POST',
			body: user,
		})
		.then(async(res) => {
			const result = await res.json()
			if (res.ok){
				return result
			}
			throw new Error(result.message || '请求失败')
		})
		.then(res => {
			dispatch(updateUserProfile(res.body))
			navigate("/chat")
		})
		.catch(err => {
			console.error(err)
		})
	}

	return (
		<section>
			<form action=''>
				<label htmlFor='account'>account:
					<input type='text' name='account' id='account' value={ account }
								 onChange={ (e) => setAccount(e.target.value) } />
				</label>password:
				<label htmlFor='password'>
					<input type='text' name='password' id='password' value={ password }
								 onChange={ (e) => setPassword(e.target.value) } />
				</label>
				<button onClick={ (event) => {
					event.preventDefault()
					submit(account, password)
				} }>Submit
				</button>
			</form>
		</section>
	)
};
