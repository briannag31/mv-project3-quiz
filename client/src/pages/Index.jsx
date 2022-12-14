import React, {useState} from "react";
import ReactCardFlip from 'react-card-flip';

const url = "http://localhost:3001/api"
 
const Card = ({project,fetchCards})=>{
  const [isFlipped, setIsFlipped] = useState(false)
  const updateKnown = async () =>{
    await fetch (`${url}/cardKnown/${project._id}`,{
      method: "PUT",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      }
    })
    fetchCards()
  }
  const undoKnown = async () =>{
    await fetch (`${url}/undoCardKnown/${project._id}`,{
      method: "PUT",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded"
      }
    })
    fetchCards()
  }
  const deleteCard = async (id) =>{
    await fetch(`${url}/deleteCards/${project._id}`, {
       method: "DELETE",
       header: {
         "Content-Type": "application/x-www-form-urlencoded"  
       },
     })
     // const data = await res.json()
     fetchCards()
   }

  return(
    <ReactCardFlip isFlipped={isFlipped} flipDirection="horizontal" >
     
        <div className="CardFront" onClick={()=> {setIsFlipped((prev)=> !prev)}}>
        <button onClick={()=> deleteCard()} type="button" class="btn-close  pull-right close-button" aria-label="Close"></button>
          <h3 className="cardFont">{project.term}</h3>
        </div>
      <div className="CardBack   align-column" >
        
          <h3 className="cardFont">{project.answer}</h3>
          <button type="button" class="btn-close  pull-right close-button-back" aria-label="Close"></button>
          <div style={{ margin: "0 auto"}} onClick={() => setIsFlipped((prev) => !prev)} >
           <p style={{marginTop:" 120px", marginBottom: "-20px"}}> Did you know the answer?</p> <br />
           <div className="card-footer">
            <div className="row">
              <div className="col-6 pr-0">
              <button onClick={()=> {updateKnown() }} type="button" className="btn btn-success"> <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-check2" viewBox="0 0 16 16">
  <path d="M13.854 3.646a.5.5 0 0 1 0 .708l-7 7a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.708L6.5 10.293l6.646-6.647a.5.5 0 0 1 .708 0z"/>
</svg>&nbsp;&nbsp;&nbsp;Yes</button>
              </div>
              <div className="col-6 pr-0">
              <button onClick={()=> {undoKnown();}} type="button" className="btn btn-danger"> <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-x-lg" viewBox="0 0 16 16">
  <path d="M2.146 2.854a.5.5 0 1 1 .708-.708L8 7.293l5.146-5.147a.5.5 0 0 1 .708.708L8.707 8l5.147 5.146a.5.5 0 0 1-.708.708L8 8.707l-5.146 5.147a.5.5 0 0 1-.708-.708L7.293 8 2.146 2.854Z"/>
</svg>&nbsp;&nbsp;&nbsp; No</button>
              </div>
            </div>
           </div>
          </div>
    </div>
    </ReactCardFlip>
  )
}

const Index = ({setViewCards,cards, setCards, setAddCard,card,setCard,newCard, setNewCard, addCard, fetchCards}) =>{
  return(
   
    <div className="row hello" >
      {cards.map((item, idx)=>{
      return  item.known ? 
          
          <div className="plainCard col-3" style={{borderColor: "green", margin: "40px", marginTop: "100px"}} >
          <Card project={item} key={idx}  fetchCards={fetchCards} />
         </div>
         :
         <div className="plainCard col-3" style={{borderColor: "red", margin: "40px", marginTop: "100px"}} >
         <Card project={item} key={idx}   fetchCards={fetchCards}/>
        </div>
      })}
    </div>
  )
}

export default Index
