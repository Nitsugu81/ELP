//Bibliothèques importées//
import inquirer from 'inquirer';
import chalk from 'chalk';
import figlet from 'figlet';
import clear from'clear';
import { exec } from 'node:child_process';
import psList from 'ps-list';
import readline from 'readline';
import shell from 'shelljs';


//CLI
async function getCommand() {

    
    const ans = await inquirer.prompt([
        {
            type: 'input',
            name: 'command',

            message: chalk.red('>>>'),
            prefix : chalk.blue(`${process.cwd()}`), //affiche le répertoire courant
        },
    ]).then(answers => {
        chooseCommand(answers.command.split(' '));
    });
    return ans;
}



//Méthode choix commande
async function chooseCommand(c) {
    
    switch (c[0]) {
        case 'exec' :
            //:`${process.cwd()}`+ "/" +
            if (c[2] == '!'){
                exec(c[1] + " &", (err, stdout, stderr) => { //En chemin absolu

                    console.log(stdout)
                    if (err) {
                        console.error(`exec error: ${err}`);
                        return;
                    }
                });
            }

            else{
                exec(c[1], (err, stdout, stderr) => {
                if (err) {
                    console.error(`exec error: ${err}`);
                    return;
                }
                console.log(stdout)
            })}
            break;

        case 'cd' :
            cd(c[1])
            break;
        case 'lp' :
            let liste = (await psList())
            for (let i = 0; i < liste.length; i++) {
                console.log(liste[i].pid + ' ' + liste[i].name);
            }
            //console.log(await psList());
            break;
        case 'end' :
            process.exit();
            break;
        case 'bing' :
            switch (c[1]) {
                case '-k' :
                    exec(`kill ${c[2]}`, (error, stdout, stderr) => {
                        if (error) {
                            console.error(`Error: ${error}`);
                            return;
                        }
                    });
                    break;
                case'-p' :
                    exec(`kill -STOP ${c[2]}`, (error, stdout, stderr) => {
                        if (error) {
                            console.error(`Error: ${error}`);
                            return;
                        }
                    });
                    break;
                case '-c' :
                    exec(`kill -CONT ${c[2]}`, (error, stdout, stderr) => {
                        if (error) {
                            console.error(`Error: ${error}`);
                            return;
                        }
                    });
                    break;
                default :
                    console.log("Invalid command. Use -k, -p, or -c.");
                    console.log(c[1])
                    break;
            }
            break;
        case 'clear' :
            clear();
            break;

        case 'mv' :
            shell.mv(c[1], c[2]);
            break;

        case 'help' :
            console.log("cd <rep> : permet de naviguer dans les fichier");
            console.log("clear : nettoie le terminal");
            console.log("lp : affiche les processus en cours");
            console.log("mv <fichier> <rep> : change le fichier de répertoire");
            console.log("crt+P : ferme le terminal");
            console.log("bing <-k|-c|-p> : Tue|Relance|Met en pause un process");
            console.log("exec <programme> : l'exécute avec un chemin absolu");
            console.log("exec <programme> ! : le lance en fond")
            break;
        default :
            console.log("La commande n'est pas reconnue");
    
        
        
    }
}


//Run le terminal
clear();
ctrl_P();
console.log(
    chalk.red(
        figlet.textSync('CLI : Mathias et Augustin', { horizontalLayout: 'full' })
    )
);
run()

async function run() {
    const command = await getCommand();
    run();
}

function ctrl_P(){
    readline.emitKeypressEvents(process.stdin);
    if (process.stdin.isTTY) process.stdin.setRawMode(true);
    process.stdin.on("keypress", (str, key) => {
    if(key.ctrl && key.name == "p") process.exit()
    })
}


function cd(path) {
    try {
        process.chdir(path);
    } catch (err) {
        console.error(`cd: ${err}`);
    }
}
