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

app.listen(3000, () => {
    console.log("Start on port 3000.")
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

app.post("/data", (req: express.Request, res: express.Response) => {
    console.log(req.body.name)
    res.sendStatus(200)
})

app.get("/takemed", async (req: express.Request, res: express.Response) => {
    const id: string = req.query.id as string
    console.log("id value = " + id + " type = " + typeof id)

    const message = await DbInquire.getMessage(parseInt(id))
    res.send(message)
})
