function getEquipment(items: Item[]): Equipment {
    let equipment: Equipment = {}
    items.forEach(el => {
        if (el.ItemState === "equipped") {
            switch (el.Category) {
                case ("helmet"):
                    equipment.helmet = el;
                    break;
                case ("chestplate"): 
                    equipment.chestplate = el;
                    break;
                case ("leggings"):
                    equipment.leggings = el;
                    break;
                case ("boots"):
                    equipment.boots = el;
                    break;
                case ("pet"):
                    equipment.pet = el;
                    break;
            }
        }
    })
    return equipment;
}
  


function getHeroHtml(e: Equipment): HTMLElement {
    let result: HTMLElement = document.createElement("div");
    let avatar: HTMLElement = document.createElement("img");
    result.setAttribute("class", "hero");
    avatar.setAttribute("id", "avatar");
    avatar.setAttribute("src", "https://static.wikia.nocookie.net/minecraft_ru_gamepedia/images/3/33/Стив_JE2_BE1.png");
    avatar.setAttribute("alt", "avatar");
    for (let key of Object.keys(e)) {
        if (e[key]) {
            let tmp = document.createElement("img")
            tmp.setAttribute("src", e[key].ImageForHero);
            tmp.setAttribute("id", e[key].Category);
            result.appendChild(tmp);
        }
    }
    result.appendChild(avatar);
    return result;
}

function renderHero(parentElementId: string): void {
    sendRequest("GET", server + "/users/3")
        .then(r => {
            let equipment: Equipment = getEquipment(r.Items);
            const parentElement: HTMLElement = document.getElementById(parentElementId) as HTMLElement;
            if (parentElement != null) {
                parentElement.prepend(getHeroHtml(equipment));    
            }
        })
        .catch(e => {
            throw e;
        }) 
}
