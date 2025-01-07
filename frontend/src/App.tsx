import {useState} from 'react'
import {axios} from "./util/axios.ts";

function App() {
    const [count, setCount] = useState(0)
    const [bar, setBar] = useState(0)

    const getFoo = () => {
        axios.get('/sample/foo')
            .then((res) => {
                setBar(res.data.bar)
            })
            .catch((err) => {
                alert(err)
            })
    }
    return (
        <div style={{display: 'flex', flexDirection: 'column', alignItems: 'start'}}>
            <button onClick={() => setCount((count) => count + 1)}>
                count is {count}
            </button>
            <button onClick={getFoo}>
                foo
            </button>

            {
                bar > 0 && (
                    <div>
                        bar is {bar}
                    </div>
                )
            }
        </div>
    )
}

export default App
