<script lang="ts">
    import {
        Navbar,
        Nav,
        NavItem,
        NavLink,
        Dropdown,
        DropdownToggle,
        DropdownMenu,
        DropdownItem,
    } from 'sveltestrap';


    function quit() {
		globalThis.Quit()
	}

	function connect() {
		globalThis.Connect()
	}

	function disconnect(e: MouseEvent) {
		globalThis.Disconnect()
	}

</script>

<Navbar color="dark" dark fixed="top" expand container="fluid">
    <Nav navbar>
        <Dropdown nav inNavbar>
            <DropdownToggle nav caret>File</DropdownToggle>
            <DropdownMenu end>
            <DropdownItem on:click={connect}>Options</DropdownItem>
            <DropdownItem on:click={disconnect}>Logs</DropdownItem>
            <DropdownItem divider />
            <DropdownItem on:click={quit}>Exit</DropdownItem>
            </DropdownMenu>
        </Dropdown>
        <NavItem>
            <NavLink>Help</NavLink>
        </NavItem>
    </Nav>
</Navbar>