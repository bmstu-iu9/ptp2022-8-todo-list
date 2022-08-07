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

function createEquipHtml(item: Item): HTMLElement {
    let html = document.createElement("img");
    html.setAttribute("class", "img-fluid");
    html.setAttribute("src", item.ImageForHero);
    html.setAttribute("id", item.Category);
    return html;
}
  

function setEquipmentImg(id: number, category: "string"): void{
    let equip: HTMLElement = document.getElementById(category)!;
    if (equip != null) {
        equip.setAttribute("src", itemsInventory[id].ImageForHero);
    } else {
        let hero: HTMLElement = document.getElementById("hero")!;
        hero.appendChild(createEquipHtml(itemsInventory[id]));
    }
}


function removeEquipmentImg(category: Category) {
    document.getElementById(category)?.remove();
}



function getHeroHtml(e: Equipment): HTMLElement {
    let result: HTMLElement = document.createElement("div");
    let avatar: HTMLElement = document.createElement("img");
    result.setAttribute("class", "hero");
    avatar.setAttribute("class", "img-fluid");
    avatar.setAttribute("id", "avatar");
    avatar.setAttribute("src", "https://static.wikia.nocookie.net/minecraft_ru_gamepedia/images/3/33/Стив_JE2_BE1.png");
    avatar.setAttribute("alt", "avatar");
    for (let key of Object.keys(e)) {
        if (e[key]) {
            result.appendChild(createEquipHtml(e[key]));
        }
    }
    result.appendChild(avatar);
    return result;
}

function renderHero(parentElementId: string, items: Item[]): void {
            let equipment: Equipment = getEquipment(items);
            const parentElement: HTMLElement = document.getElementById(parentElementId) as HTMLElement;
            if (parentElement != null) {
                parentElement.prepend(getHeroHtml(equipment));    
            }
 
}
