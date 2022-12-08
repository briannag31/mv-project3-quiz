import React, { useState, useEffect } from "react";
import Header from "./components/Header"
import Footer from "./components/Footer"
import Nav from "./components/Nav"
import New from "./pages/New"
import Index from "./pages/Index"


function App() {
  const [cards, setCards] = useState([])
  const [card, setCard] = useState(null)
  const [addCard, setAddCard] = useState(false)
  const [newCard, setNewCard] = useState({
    collection: "",
    term:"",
    answer: "",
    known: false
  })


  const url = "http://localhost:3001/api"

  async function fetchCards(){
    try{
      const res = await fetch(`${url}/cards`)
      const data = await res.json()
      setCards(data)
    } catch(err){
      console.log(err)
    }
  }

  useEffect(() =>{
    fetchCards()
  }, [])
  return (

    <div className="App">
      <Nav />
      <Header />
      {addCard ?
      
      <New  
        cards={cards}
        setCards={setCards}
        addCard={addCard}
        setAddCard={setAddCard}
        card={card}
        setCard={setCard}
        newCard={newCard}
        setNewCard={setNewCard}
        fetchCards={fetchCards}
        /> 
        :
        <Index 
        fetchCards={fetchCards}
        cards={cards}
        setCards={setCards}
        addCard={addCard}
        setAddCard={setAddCard}
        card={card}
        setCard={setCard}
        newCard={newCard}
        setNewCard={setNewCard}
        />
      }
        <Footer />
    </div>
  );
}

export default App;
