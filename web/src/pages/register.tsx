import { useState } from 'react'
import '../assets/style/register.css'
import { useNavigate } from 'react-router-dom'

export default function Register () {
	const [account, setAccount] = useState<string>('sun')
	const [password, setPassword] = useState<string>('123')
	const [avatar, setAvatar] = useState<string>('avatar')
	const [email, setEmail] = useState<string>('xicon@qq.com')
	const [nickname, setNickname] = useState<string>('s')
	const [gender, setGender] = useState<string>('girl')
	const navigate = useNavigate()

	const submit = () => {
		const userInfo: FormData = new FormData()
		userInfo.append('account', account)
		userInfo.append('password', password)
		userInfo.append('avatar', avatar)
		userInfo.append('email', email)
		userInfo.append('nickname', nickname)
		userInfo.append('gender', gender)
		userInfo.append('createdAt', new Date().getTime()
		.toString())
		userInfo.append('updatedAt', new Date().getTime()
		.toString())

		fetch('http://localhost:4000/register', {
			method: 'PUT',
			body: userInfo,
		})
		.then(async(res) => {
			const body = await res.json()
			if (body.ok){
				return body
			}
			else if (body.code === 403){
				navigate('/register')
			}
			throw new Error('请求失败')
		})
		.then((res) => {
			console.log(res)
		})
		.catch((err) => {
			console.error(err)
		})
	}

	return (
		<>
			<form>
				<label htmlFor='account'>account:
					<input type='text' name='account' id='account' value={ account }
								 onChange={ (e) => setAccount(e.target.value) } />
				</label>

				<label htmlFor='password'>password:
					<input type='text' name='password' id='password' value={ password }
								 onChange={ (e) => setPassword(e.target.value) } />
				</label>

				<label htmlFor='email'>email:
					<input type='text' name='email' id='email' value={ email } onChange={ (e) => setEmail(e.target.value) } />
				</label>

				<label htmlFor='avatar'>avatar:
					<input type='text' name='avatar' id='avatar' value={ avatar } onChange={ (e) => setAvatar(e.target.value) } />
				</label>

				<label htmlFor='nickname'>nickname:
					<input type='text' name='nickname' id='nickname' value={ nickname }
								 onChange={ (e) => setNickname(e.target.value) } />
				</label>

				<label htmlFor='gender'>gender:
					<input type='text' name='gender' id='gender' value={ gender } onChange={ (e) => setGender(e.target.value) } />
				</label>

				<button type={ 'button' } onClick={ submit }>Submit</button>
			</form>
		</>
	)
};
