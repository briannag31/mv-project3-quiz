
export default function New({cards, setCards, setAddCard,card,setCard,newCard, setNewCard, fetchCards}){
    const handleChange = e =>{
        const updatedCard = {
            ...newCard, [e.target.name]:e.target.value
        }
        setNewCard(updatedCard)
    }
    const url = "http://localhost:3001/api"

    const handleCancel = async()=>{
        // const res = await(url)
        // const data = await res.json()
        setAddCard(false)
    }

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
    <form onSubmit={handleSubmit}>
        <div>
            <div> 
                <input type="text" placeholder="Collection Name" name="collection" value={newCard.collection} onChange={handleChange}/>
            </div>  
            <div> 
                <input type="text" placeholder="Vocab Term" name="term" value={newCard.term} onChange={handleChange}/>
            </div>
            <div>
            <input type="text"  placeholder="Answer" name="answer" value={newCard.answer} onChange={handleChange}/>

            </div>
            <div> 
                <p>Do you know this term?</p>
                <label>
                    <input  type="radio" value={newCard.known} checked={true} onChange={handleChange}/>
                    no
                </label>
            </div>
            <div> 
            <label>
                    <input  type="radio" value={newCard.known}  onChange={handleChange}/>
                    yes
                </label>
            </div>
            <button type="submit" className="btn btn-success">Add New Card</button>
        </div>
    </form>
    <button style={{width: "100%", marginTop: "20px"}} className="btn btn-danger" onClick={() => handleCancel()}>Cancel</button>
    </>

}