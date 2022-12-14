
export default function New({cards, setCards, setAddCard,card,setCard,newCard, setNewCard, fetchCards}){
    const handleChange = e =>{
        const updatedCard = {
            ...newCard, [e.target.name]:e.target.value
        }
        setNewCard(updatedCard)
    }
    const url = "http://localhost:3001/api"

   

    const createCard = async(cardData)=>{
        try{
            await fetch(`${url}/cards`, {
                method: "POST", 
                headers: {
                  "Content-Type": "application/x-www-form-urlencoded"  
                },
                body: JSON.stringify(cardData)
            })
            fetchCards();
        } catch(error){
            console.log(error);
        }
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        createCard(newCard)
        setNewCard({
            collection: "",
            term:"",
            answer: "",
            known: false
        })
        setAddCard(false);
        console.log("new set of cards",cards)
    }

    return<>
    <h1 className="form-title">Add A New Card</h1>
    <form onSubmit={handleSubmit}>
        <div className="form-group">
            <div className="form-floating"> 
                <input className="form-control" id="floatingTerm" type="text" placeholder="Term" name="term" value={newCard.term} onChange={handleChange}/>
                <label for="floatingTerm">Term</label>
            </div>
            <div className="form-floating ans">
            <input className="form-control" id="floatingAnswer" type="text"  placeholder="Answer" name="answer" value={newCard.answer} onChange={handleChange}/>
            <label for="floatingAnswer">Answer</label>
            </div>
            <button style={{width:"100%"}} type="submit" className="btn btn-success">Add New Card</button>
        </div>

    </form>
    <button style={{width:"100%", marginTop: "20px"}}  className="btn btn-danger" onClick={() => setAddCard(false)}>Cancel</button>
    </>

}