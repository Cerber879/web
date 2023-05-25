let user = {
    "Email": null,
    "Password": null
}

function pass_visibility()
{
    pass = document.getElementById('pass_field');

    if(pass.type === 'text')
    {
        pass.type = 'password';
    }
    else if(pass.type === 'password')
    {
        pass.type = 'text';
    }
}

function Select(type)
{
    let field = document.getElementById(type + "_field");

    if(!field.classList.contains('main__field_error'))
    {
        field.classList.add('main__field_select');   
    }
}

function NotSelect(type)
{
    let field = document.getElementById(type + "_field");
        block = document.getElementById(type + "_block");
    
    field.classList.remove('main__field_select')

    if((field.value === "") && (!field.classList.contains('main__field_error')))
    {
        field.classList.add('main__field_error');

        let label = document.createElement('p');

        switch (type) {
            case 'login':
                label.textContent = "Email is required.";
                break;
            case 'pass':
                label.textContent = "Password is required.";
                break;
        }

        label.classList.add('main__error');

        block.insertBefore(label, block.children[2]);
    }
    else if(!(field.value === ""))
    {
        if(field.classList.contains('main__field_error'))
        {
            field.classList.remove('main__field_error');
            block.children[2].remove();
        }

        switch (type) {
            case 'login':
                user.Email = field.value;
                break;
            case 'pass':
                user.Password = field.value;
                break;
        }
    }
}

function DataError()
{
    let massage = document.getElementById('massage');
    if(!massage.classList.contains('main__massage'))
    {
        let icon = document.createElement('img');
        icon.classList.add('main__icon');
        icon.src = "../static/svg_files/alert_circle.svg";
        
        let text = document.createElement('p');
        text.classList.add('main__massage-text');
        text.textContent = "A-Ah! Check all fields";
    
        massage.classList.add('main__massage');
        massage.insertBefore(text, massage.children[0]);
        massage.insertBefore(icon, massage.children[0]);
    }
}

function Click()
{
    if((user.Email === null) && (user.Password === null))
    {
        DataError();
        NotSelect('login');
        NotSelect('pass');
    }
    else
    {
        let email = user.Email;
        let email_name = "";
        let email_domen_name = "";
        let email_domen = "";
        let part_of_email = 1;
        let is_email_valid = true;
    
        for(let i = 0; i < email.length; i++)
        {
            let ch = email[i];
            if(((ch >= 'A') && (ch <= 'Z')) || ((ch >= 'a') && (ch <= 'z')) || ((ch == "@") && (part_of_email == 1)) || ((ch == ".") && (part_of_email == 2)))
            {
                if((ch == "@") && (part_of_email == 1))
                {
                    part_of_email = 2
                }
                else if((part_of_email == 2) && (ch == "."))
                {
                    part_of_email = 3
                }
                else if(part_of_email == 1)
                {
                    email_name += ch;
                }
                else if(part_of_email == 2)
                {
                    email_domen_name += ch;
                }
                else if(part_of_email == 3)
                {
                    email_domen += ch;
                }
            }
            else
            {
                is_email_valid = false;
            }
        }
        
        if(!((email_name !== "") && (email_domen_name !== "") && (email_domen !== "")) || !is_email_valid)
        {
            is_email_valid = false;
    
            let field = document.getElementById('login_field');
            if(!field.classList.contains("main__field_error"))
            {
                field.classList.add('main__field_error');
                
                let err_message = document.createElement('p');
                err_message.classList.add('main__error');
                err_message.textContent = "Incorrect email format. Correct format is ****@**.***";
        
                let block = document.getElementById('login_block');
                block.insertBefore(err_message, block.children[2]);
            }

            DataError();
        }
            
        if(is_email_valid)
        {
            window.location.replace("http://localhost:3000/admin")
        }
    }
}