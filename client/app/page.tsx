'use client'
import { ChangeEvent, useCallback, useState } from 'react'

interface Person {
	username: string
	password: string
}

interface UserInfo {
	msg: string
	tokenString: string

	person: Person
	code: number
}

interface Error {
	msg: string
	code: number
}

type Result = UserInfo | Error

const useUserInfo = (usr: string, pwd: string) => {
	const [result, setResult] = useState<Result | Error>({ code: 200, msg: '', person: { username: '', password: '' }, tokenString: '' })
	const get = useCallback(() => {
		fetch('http://localhost:4000/user/auth', {
			method: 'POST',
			headers: { 'Content-Type': 'application' },
			body: JSON.stringify({
				username: usr,
				password: pwd,
			}),
		})
		.then(async res => {
			if (!res.ok){
				setResult({ code: 400, msg: '请求异常' } as Error)
			}
			return res.json()
		})
		.then((res: Result) => {
			console.log(res)
			setResult(res as Result)
		})
		.catch((err) => {
			setResult(err)
		})

	}, [usr, pwd])

	return [result, get]
}

type JWT = Omit<Result, 'person'> & { result: Omit<Person, 'password'> }
const jwtFetch = (token: string, setJwtResult: (result: JWT | Error) => void) => {
	fetch('http://localhost:4000/user/', {
		method: 'GET',
		headers: { 'Authorization': `Bearer ${ token }` },
	})
	.then(res => {
		if (!res.ok){
			throw new Error('ERROR')
		}
		return res.json()
	})
	.then(res => {
		console.log(res)
		setJwtResult(res as JWT)

	})
	.catch(err => {
		console.error(err)
	})
}

export default function Home () {
	const [usr, setUsr] = useState<string>('lisa')
	const [pwd, setPwd] = useState<string>('123456')
	const [result, setResult] = useUserInfo(usr, pwd)
	const [jwtResult, setJwtResult] = useState<JWT>({ code: 200, msg: '', result: { username: '' } })

	return (
		<section>
			{
				result.code == 200
					? (<ol>
						<li>
							<p>返回的消息将显示在这里:{ (result as UserInfo).msg || 'msg' }</p>

						</li>
						<li>
							<p>token将显示在这里:{ (result as UserInfo).tokenString || 'token' }</p>

						</li>
						<li><p>用户账号将显示在这里:{ (result as UserInfo).person.username || 'username' }</p></li>
						<li><p>用户密码将显示在这里:{ (result as UserInfo).person.password || 'password' }</p></li>
					</ol>)
					: (<aside>
						<ol>
							<li>
								<p>错误信息:{ result.msg || '错误信息' }</p>
								<p>状态码:{ result.code || '状态码' }</p>
							</li>
						</ol>
					</aside>)
			}
			<aside>
				<label htmlFor=''>
					<input
						placeholder='Name'
						value={ usr }
						onChange={ (e: ChangeEvent<HTMLInputElement>) => setUsr(e.currentTarget.value) } />
				</label>
				<label htmlFor=''>
					<input
						type='password'
						placeholder='Password'
						value={ pwd }
						onChange={ (e: ChangeEvent<HTMLInputElement>) => setPwd(e.currentTarget.value) } /></label>
				<button
					color={ 'gradient' }
					onClick={ () => setResult(usr, pwd) }
				>Submit
				</button>
			</aside>

			<aside>
				<section>
					<ol>
						<li><p>msg:{ jwtResult.msg }</p></li>
						<li><p>code:{ jwtResult.code }</p></li>
						<li><p>username:{ jwtResult.result.username }</p></li>
						<li><p>tokenString:{ jwtResult.tokenString }</p></li>
					</ol>
				</section>

				<label htmlFor=''>
					<input
						placeholder='input token'
						value={ 'tokenString' in result ? result.tokenString : '' }
						onChange={ (e: ChangeEvent<HTMLInputElement>) => setUsr(e.currentTarget.value) } /></label>
				<button
					color={ 'gradient' }
					onClick={ () => jwtFetch(result.tokenString, setJwtResult) }
				>Submit
				</button>
			</aside>
		</section>
	)
}
