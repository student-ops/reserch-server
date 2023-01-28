import express from "express"
const app: express.Express = express()
app.use(express.json())
app.use(express.urlencoded({ extended: true }))
import * as DbInquire from "./db-inquire"

//CROS対応（というか完全無防備：本番環境ではだめ絶対）
app.use(
    (
        req: express.Request,
        res: express.Response,
        next: express.NextFunction
    ) => {
        res.header("Access-Control-Allow-Origin", "*")
        res.header("Access-Control-Allow-Methods", "*")
        res.header("Access-Control-Allow-Headers", "*")
        next()
    }
)

const port: number = 8080
app.listen(port, () => {
    console.log("Start on port" + port)
})

type User = {
    id: number
    name: string
    email: string
}

type Taken = {
    time: number
    medicine: string
}

app.get("/ping", async (req: express.Request, res: express.Response) => {
    console.log("reached ping")
    res.sendStatus(200)
})

app.post("/json", async (req: express.Request, res: express.Response) => {
    console.log(req.body.id)
    res.sendStatus(200)
})

app.post("/takemed", async (req: express.Request, res: express.Response) => {
    var id: string = req.body.id
    const message = await DbInquire.getMessage(parseInt(id))
    type TakemedRespones = {
        message: string
    }
    var respmessage: TakemedRespones = {
        message: message,
    }
    res.status(200).send(respmessage)
})
